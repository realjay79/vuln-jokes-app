import React, { useState } from "react";
import "./App.css";

function App() {
  const [jokeId, setJokeId] = useState("");
  const [jokeResult, setJokeResult] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await fetch(`http://localhost:8080/joke?id=${jokeId}`, {
        method: "GET",
        credentials: "include", // 🍪 CSRF-prone
      });

      const data = await response.text();

      // XSS-prone: directly injecting HTML
      setJokeResult(data);
    } catch (error) {
      console.error("Error fetching joke:", error);
    }
  };

  return (
    <div className="container">
      <h1 className="main-header">Mmm... Jokes.</h1>
      <p className="sub-header">Because real therapy costs money.</p>

      <form onSubmit={handleSubmit} className="joke-form">
        <input
          type="text"
          placeholder="Wanna laugh? Type ‘random’. Or don’t. I’m not your boss."
          value={jokeId}
          onChange={(e) => setJokeId(e.target.value)}
          className="joke-input"
        />
        <button type="submit" className="joke-button">
          Fetch Joke
        </button>
      </form>

      <div
        className="joke-output"
        dangerouslySetInnerHTML={{ __html: jokeResult }}
      />
    </div>
  );
}

export default App;
