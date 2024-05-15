import { useEffect, useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import * as apiService from "../apiservice/apiservice";
import { UserContext } from "../index";

export const Signin = () => {
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
