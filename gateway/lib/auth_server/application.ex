defmodule AuthServer.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false
  require Logger
  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Starts a worker by calling: AuthServer.Worker.start_link(arg)
      # {AuthServer.Worker, arg}
      {Plug.Cowboy, scheme: :http, plug: AuthServer.HelloWorld, options: [port: 5002]},
      {Plug.Cowboy, scheme: :http, plug: AuthServer.Router, options: [port: 5003]}
    ]
    Logger.info "Running web server(s)."
    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: AuthServer.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
