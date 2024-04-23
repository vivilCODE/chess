import { GoogleLogin } from "@react-oauth/google";

export const Login = (props) => {
  return (
    <GoogleLogin
      onSuccess={(credentialResponse) => {
        console.log("ONSUCCESS: ", credentialResponse);
      }}
      onError={() => {
        console.log("Login Failed");
      }}
    />
  );
};
