export default function Cell({ value, onClick, playerColor }) {
  let className = "cell";
  let style = {};

  if (value === 1) {
    className += " drop";
    style.backgroundColor = playerColor;
  } else if (value === 2) {
    className += " drop";
    style.backgroundColor = "#FFD700"; // Bot color (gold)
  }

  return <div className={className} style={style} onClick={onClick}></div>;
}
