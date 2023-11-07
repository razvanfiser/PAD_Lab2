defmodule AuthServer.Router do
  use Plug.Router

  plug :match
  plug :dispatch

  get "/" do
    send_resp(conn, 200, "Welcome my nigg!")
  end

  match _ do
    send_resp(conn, 400, "Lol nothing!!!")
  end
end
