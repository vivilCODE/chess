import * as apiService from "../apiservice/apiservice";
import { GameWindow } from "../partials/GameWindow";
import { Login } from "../components/Login/Login";

export const LandingPage = () => {
  return (
    <div className="landing-page">
      <Login />
      <button onClick={apiService.Ping}>Ping</button>
      <GameWindow />
    </div>
  );
};
