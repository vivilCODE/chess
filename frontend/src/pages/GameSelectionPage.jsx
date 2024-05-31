import { GameSelectionWindow } from "../partials/GameSelectionWindow/GameSelectionWindow";
import { Link } from "react-router-dom";

export const GameSelectionPage = () => {
  return (
    <>
      game selection page
      <Link to={"/"}>Home</Link>
      <GameSelectionWindow />
    </>
  );
};
