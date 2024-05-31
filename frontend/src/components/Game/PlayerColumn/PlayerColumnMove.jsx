import "./move.css";

export const PlayerColumnMove = () => {
  return (
    <li className="player-column__moves-list__move">
      <div className="player-column__moves-list__move__left">
        <span className="player-column__moves-list__move__left__number">
          1.
        </span>
        <img
          className="player-column__moves-list__move__left__piece-icon"
          src="./pieces/black-pawn.svg"
          alt="Piece icon"
        />
        <span className="player-column__moves-list__move__left__move-name">
          e4
        </span>
      </div>
      <div className="player-column__moves-list__move__right">
        <span className="player-column__moves-list__move__right__opponent-move">
          f6
        </span>
      </div>
    </li>
  );
};
