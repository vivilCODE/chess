import { Link } from "react-router-dom";

export const SignInButton = () => {
  let loginUrl = `https://accounts.google.com/signin/oauth?response_type=code&client_id=${process.env.REACT_APP_GAPI_CLIENT_ID}&scope=openid%20email&redirect_uri=${process.env.REACT_APP_REDIRECT_URL}&state=STATE&nonce=NONCE `;

  return (
    <li className="navigation__list__sign-in">
      <Link to={loginUrl}>Sign in with Google</Link>
    </li>
  );
};
