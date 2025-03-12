import { useState } from "react";

const LoginForm: React.FC = () => {
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const [messages, setMessages] = useState<string[]>([]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    if (name === "userName") setUserName(value);
    if (name === "password") setPassword(value);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const errors = [];
    if (!userName) errors.push("Username is required.");
    if (!password) errors.push("Password is required.");
    setMessages(errors);

    if (errors.length === 0) {
      console.log("Logging in:", { userName, password });
    }
  };

  return (
    <div className="flex justify-center items-center h-screen bg-gray-200">
      <div className="bg-white p-6 rounded-lg w-80 shadow-lg">
        <h2 className="text-xl font-bold mb-4">Login</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block mb-1 font-medium" htmlFor="userName">
              User Name:
            </label>
            <input
              type="text"
              id="userName"
              name="userName"
              value={userName}
              onChange={handleChange}
              required
              className="w-full p-2 border border-gray-300 rounded"
            />
          </div>
          <div className="mb-4">
            <label className="block mb-1 font-medium" htmlFor="password">
              Password:
            </label>
            <input
              type="password"
              id="password"
              name="password"
              value={password}
              onChange={handleChange}
              required
              className="w-full p-2 border border-gray-300 rounded"
            />
          </div>

          {messages.length > 0 && (
            <div className="mb-4">
              {messages.map((msg, index) => (
                <p key={index} className="text-red-500 text-sm">
                  {msg}
                </p>
              ))}
            </div>
          )}

          <div className="mb-4 flex justify-center">
            <button
              type="submit"
              className="w-1/2 p-2 bg-green-500 text-white text-lg rounded cursor-pointer hover:bg-green-600"
            >
              Login
            </button>
          </div>

          <div className="text-center">
            Don't have an account?{" "}
            <a href="/register" className="text-blue-500 hover:underline">
              Register
            </a>
          </div>
        </form>
      </div>
    </div>
  );
};

export default LoginForm;
