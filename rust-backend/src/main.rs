use crate::game::{Game, GameInputMessage, GameSender};
use crate::server::client::Client;
use simplelog::*;
use std::sync::atomic::{AtomicU32, Ordering};
use tokio::sync::mpsc;
use tokio_stream::wrappers::UnboundedReceiverStream;
use tracing::{debug, error};
use warp::filters::ws::WebSocket;
use warp::Filter;

pub mod game;
pub mod server;
pub mod superstellar;
pub mod types;

/// Our global unique user id counter.
static NEXT_USER_ID: AtomicU32 = AtomicU32::new(1);

#[tokio::main]
async fn main() {
    CombinedLogger::init(vec![TermLogger::new(
        LevelFilter::Debug,
        Config::default(),
        TerminalMode::Mixed,
    )])
    .unwrap();

    let (game_sender, game_receiver) = mpsc::unbounded_channel();
    let game_receiver = UnboundedReceiverStream::new(game_receiver);
    let game_sender_filter = warp::any().map(move || game_sender.clone());

    let routes = warp::path!("superstellar")
        .and(warp::ws())
        .and(game_sender_filter)
        .map(|ws: warp::ws::Ws, game_sender: GameSender| {
            ws.on_upgrade(move |websocket| user_connected(websocket, game_sender))
        });

    tokio::spawn(async move {
        let mut game = Game::new(game_receiver);

        game.run().await;
    });

    warp::serve(routes).run(([127, 0, 0, 1], 8081)).await;
}

async fn user_connected(websocket: WebSocket, game_sender: GameSender) {
    let (game_output_sender, game_output_receiver) = mpsc::unbounded_channel();
    let game_output_receiver = UnboundedReceiverStream::new(game_output_receiver);
    let id = NEXT_USER_ID.fetch_add(1, Ordering::Relaxed);
    let mut client = Client::new(id, websocket, game_sender.clone(), game_output_receiver);

    game_sender
        .send(GameInputMessage::PlayerConnected {
            id: client.id,
            sender: game_output_sender,
        })
        .expect("Can't send message to game");

    debug!("User {} connected", client.id);

    match client.run().await {
        Ok(()) => {}
        Err(inner) => {
            error!("Client error: {:?}", inner);
        }
    };
}
