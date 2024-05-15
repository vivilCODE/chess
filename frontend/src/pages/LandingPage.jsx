import * as apiService from "../apiservice/apiservice";
import { GameWindow } from "../partials/GameWindow";
import { Login } from "../components/Login/Login";
import { UserContext } from "../index";
import { useContext, useEffect } from "react";

export const LandingPage = () => {
  const [user, setUser] = useContext(UserContext);

  useEffect(() => {
    console.log("user: ", user);
  }, [user]);

  return (
    <div className="landing-page">
      <Login />

      <p>{user ? "signed in as " + user.user.name : "not signed in"}</p>

      <button onClick={apiService.Ping}>Ping</button>
      <GameWindow />
    </div>
  );
};
