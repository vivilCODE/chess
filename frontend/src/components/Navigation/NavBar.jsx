import { useContext } from "react";
import { Link } from "react-router-dom";
import { UserContext } from "../../index";
import { SignInButton } from "./SignInButton";

import "./navbar.css";

export const NavBar = () => {
  const [user] = useContext(UserContext);

  return (
    <div className="navigation">
      <Link to={"/"}>Home</Link>
      <ul className="navigation__list">
        {user ? <li>{user.user.name}</li> : <SignInButton />}
      </ul>
    </div>
  );
};
