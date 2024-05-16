import { Move } from "../Move/Move";

import "./playercolumn.css";

export const PlayerColumn = () => {
  return (
    <div className="player-column">
      <div className="player-column__player-info">
        <img
          className="player-column__player-info__profile-picture"
          src="face.jpg"
          alt="Profile picture"
        />
        <span className="player-column__player-info__name">EMMY BÄCKSTRÖM</span>
        <span className="player-column__player-info__color">White</span>
      </div>
      <ul className="player-column__moves-list">
        <Move />
        <Move />
        <Move />
        <Move />
        <Move />
      </ul>
      <span className="player-column__timer">05:29</span>
      <button className="player-column__resign-btn">RESIGN</button>
    </div>
  );
};
