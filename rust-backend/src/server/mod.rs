mod client;

use crate::server::client::Client;
use std::collections::HashMap;
use std::convert::Infallible;
use std::net::SocketAddr;
use std::sync::Arc;

use futures::{FutureExt, StreamExt};
use tokio::sync::mpsc::UnboundedSender;
use tokio::sync::{mpsc, RwLock};
use tokio_stream::wrappers::UnboundedReceiverStream;
use tracing::debug;
use uuid::Uuid;
use warp::filters::ws::WebSocket;
use warp::ws::Message;
use warp::Filter;

struct ClientChannels {
    sender: UnboundedSender<Result<Message, warp::Error>>,
    receiver: UnboundedReceiverStream<Result<Message, warp::Error>>,
}

pub struct Server {
    pub address: SocketAddr,
}

struct State {
    users: HashMap<Uuid, ClientChannels>,
}

impl State {
    pub fn new() -> State {
        State {
            users: HashMap::new(),
        }
    }

    pub fn register_client(&mut self, websocket: WebSocket) -> Client {
        let (client_sender, client_receiver) = mpsc::unbounded_channel();
        let client_receiver = UnboundedReceiverStream::new(client_receiver);
        let (server_sender, server_receiver) = mpsc::unbounded_channel();
        let server_receiver = UnboundedReceiverStream::new(server_receiver);

        let client = Client::new(websocket, server_sender, client_receiver);

        // Save the sender in our list of connected users.
        self.users.insert(
            client.id,
            ClientChannels {
                sender: client_sender,
                receiver: server_receiver,
            },
        );

        client
    }
}

impl Server {
    pub fn new(address: &'static str) -> Server {
        let address = address.parse().expect("Cant parse server IP address");

        Server { address }
    }

    pub async fn run(&self) {
        debug!("Starting server");

        let state = Arc::new(RwLock::new(State::new()));
        let state = warp::any().map(move || state.clone());

        let routes = warp::path!("superstellar").and(warp::ws()).and(state).map(
            |ws: warp::ws::Ws, state: Arc<RwLock<State>>| {
                ws.on_upgrade(move |websocket| user_connected(websocket, state))
            },
        );

        warp::serve(routes).run(self.address).await;
    }
}

async fn user_connected(websocket: WebSocket, state: Arc<RwLock<State>>) {
    let mut client = state.write().await.register_client(websocket);
    debug!("User {} connected", client.id);

    client.run().await
}
