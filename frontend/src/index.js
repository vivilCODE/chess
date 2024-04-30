import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { GoogleOAuthProvider } from "@react-oauth/google";

// Reset all default styles
import "./styles/css-reset.css";
import { LandingPage } from "./pages/LandingPage";
import { Signin } from "./pages/Signin";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />,
  },
  {
    path: "/signin",
    element: <Signin />,
  },
]);

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  // <React.StrictMode>
  <GoogleOAuthProvider clientId={process.env.REACT_APP_GAPI_CLIENT_ID}>
    <RouterProvider router={router} />
  </GoogleOAuthProvider>
  // </React.StrictMode>
);
