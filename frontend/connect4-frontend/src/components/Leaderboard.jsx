import { useEffect, useState } from "react";

function Leaderboard({ refreshKey }) {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch("https://four-in-a-row-backend-qq63.onrender.com/leaderboard")
      .then(res => res.json())
      .then(d => {
        console.log("LEADERBOARD DATA:", d);
        setData(d);
      });
  }, [refreshKey]);

  return (
    <table style={styles.table}>
      <thead>
        <tr>
          <th style={styles.th}>Player</th>
          <th style={styles.th}>Games</th>
          <th style={styles.th}>Wins</th>
          <th style={styles.th}>Losses</th>
          <th style={styles.th}>Draws</th>
          <th style={styles.th}>Bot Games</th>
        </tr>
      </thead>

      <tbody>
        {data.map((row, index) => (
          <tr key={`${row.Username}-${index}`}>
            <td style={styles.td}>{row.Username}</td>
            <td style={styles.td}>{row.GamesPlayed}</td>
            <td style={styles.td}>{row.Wins}</td>
            <td style={styles.td}>{row.Losses}</td>
            <td style={styles.td}>{row.Draws}</td>
            <td style={styles.td}>{row.BotGames}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

const styles = {
  table: {
    width: "100%",
    borderCollapse: "collapse",
    fontSize: "16px",
    color: "#fff",
  },
  th: {
    padding: "12px",
    background: "#0a2f4f",
    borderBottom: "2px solid #1f5a8f",
    textAlign: "left",
    fontWeight: "600",
  },
  td: {
    padding: "12px",
    borderBottom: "1px solid #1f5a8f",
  },
};

export default Leaderboard;
