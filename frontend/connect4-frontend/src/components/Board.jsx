import Cell from "./Cell";

export default function Board({ board, onPlay, playerColor }) {
  return (
    <div className="board">
      {board.map((row, r) =>
        row.map((cell, c) => (
          <Cell
            key={`${r}-${c}`}
            value={cell}
            onClick={() => onPlay(c)}
            playerColor={playerColor}
          />
        ))
      )}
    </div>
  );
}
