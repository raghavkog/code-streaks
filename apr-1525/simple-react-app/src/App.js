// src/App.js
import React, { useState } from 'react';

function App() {
  const [value, setValue] = useState('');
  const [submitted, setSubmitted] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    setSubmitted(value);
    setValue('');
  };

  return (
    <div style={{ padding: '20px' }}>
      <h1>Simple React App</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={value}
          onChange={(e) => setValue(e.target.value)}
          placeholder="Type something..."
        />
        <button type="submit" style={{ marginLeft: '10px' }}>Submit</button>
      </form>
      {submitted && <p>You submitted: {submitted}</p>}
    </div>
  );
}

export default App;
