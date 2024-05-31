import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { GoogleOAuthProvider } from "@react-oauth/google";
import { useState, createContext } from "react";

// Reset all default styles
import "./styles/css-reset.css";

// Set global style rules
import "./styles/global.css";

import { LandingPage } from "./pages/LandingPage";
import { SigninPage } from "./pages/SigninPage";
import { GameSelectionPage } from "./pages/GameSelectionPage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />,
  },
  {
    path: "/signin",
    element: <SigninPage />,
  },
  {
    path: "/play",
    element: <GameSelectionPage />,
  },
]);

export const UserContext = createContext(null);

const App = () => {
  const [user, setUser] = useState(null);

  return (
    <GoogleOAuthProvider clientId={process.env.REACT_APP_GAPI_CLIENT_ID}>
      <UserContext.Provider value={[user, setUser]}>
        <RouterProvider router={router} />
      </UserContext.Provider>
    </GoogleOAuthProvider>
  );
};

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  // <React.StrictMode>
  <App />
  // </React.StrictMode>
);
