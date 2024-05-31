import { useEffect, useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import * as apiService from "../apiservice/apiservice";
import { UserContext } from "../index";

// This page serves as the redirect page where the google api automatically points the user
// after retrieving the login code.
// The sign in page should sign the user into the chessapi service, and once it has successfully retreived
// the user data it should simply redirect to the page where the client was prior to logging in.
export const SigninPage = () => {
  const [user, setUser] = useContext(UserContext);
  const navigate = useNavigate();

  const queryParameters = new URLSearchParams(window.location.search);
  const code = queryParameters.get("code");

  useEffect(() => {
    const requestSignIn = async () => {
      let res = await apiService.SignIn(code);
      setUser(() => res);
    };

    if (code) {
      requestSignIn();
    }
  }, [code]);

  useEffect(() => {
    if (user) {
      console.log("user: ", user);
      navigate("/");
    }
  }, [user]);

  return <>Signing in</>;
};
