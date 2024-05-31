import * as apiService from "../apiservice/apiservice";
import { Login } from "../components/Login/Login";
import { UserContext } from "../index";
import { useContext, useEffect } from "react";
import { Chatroom } from "../components/Chatroom/chatroom";
import { Link } from "react-router-dom";

export const LandingPage = () => {
  const [user, setUser] = useContext(UserContext);

  useEffect(() => {
    console.log("user: ", user);
  }, [user]);

  return (
    <div className="landing-page">
      <Login />

      <p>{user ? "signed in as " + user.user.name : "not signed in"}</p>

      {/* <Chatroom /> */}
      <Link to={"/play"}>Play</Link>
    </div>
  );
};
