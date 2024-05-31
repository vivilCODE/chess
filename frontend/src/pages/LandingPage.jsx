import * as apiService from "../apiservice/apiservice";
import { UserContext } from "../index";
import { useContext, useEffect } from "react";
import { Link } from "react-router-dom";
import { NavBar } from "../components/Navigation/NavBar";

export const LandingPage = () => {
  const [user, setUser] = useContext(UserContext);

  useEffect(() => {
    console.log("user: ", user);
  }, [user]);

  return (
    <div className="landing-page">
      <NavBar />

      <Link to={"/play"}>Play</Link>
    </div>
  );
};
