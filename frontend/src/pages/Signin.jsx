import { useEffect, useState } from "react";
import * as apiService from "../apiservice/apiservice";

export const Signin = () => {
  const [user, setUser] = useState(null);

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
    }
  }, [user]);

  return <>Signing in</>;
};
