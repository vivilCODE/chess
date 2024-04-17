import * as apiService from "../apiservice/apiservice";
import { GameWindow } from "../partials/GameWindow";

export const LandingPage = () => {
  return (
    <div className="landing-page">
      <button onClick={apiService.Ping}>Ping</button>
      <GameWindow />
    </div>
  );
};
