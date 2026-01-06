import { useState } from "react";

export default function UsernameModal({ onSubmit }) {
  const [name, setName] = useState("");

  return (
    <div>
      <h3>Enter Username</h3>
      <input
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <button onClick={() => onSubmit(name)}>
        Start Game
      </button>
    </div>
  );
}
