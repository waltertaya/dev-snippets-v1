import { useState } from "react";

const RegistrationForm: React.FC = () => {
  const [formData, setFormData] = useState({
    userName: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const [messages, setMessages] = useState<string[]>([]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const errors: string[] = [];

    if (formData.password !== formData.confirmPassword) {
      errors.push("Passwords do not match!");
    }

    setMessages(errors);

    if (errors.length === 0) {
      console.log("Form submitted:", formData);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-200">
      <div className="bg-white p-6 rounded-lg shadow-md w-80">
        <h2 className="text-xl font-semibold mb-4 text-center">Register</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label className="block mb-1 font-medium" htmlFor="userName">
              User Name:
            </label>
            <input
              type="text"
              id="userName"
              name="userName"
              value={formData.userName}
              onChange={handleChange}
              className="w-full p-2 border border-gray-300 rounded"
              required
            />
          </div>

          <div className="mb-3">
            <label className="block mb-1 font-medium" htmlFor="email">
              Email:
            </label>
            <input
              type="email"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              className="w-full p-2 border border-gray-300 rounded"
              required
            />
          </div>

          <div className="mb-3">
            <label className="block mb-1 font-medium" htmlFor="password">
              Password:
            </label>
            <input
              type="password"
              id="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
              className="w-full p-2 border border-gray-300 rounded"
              required
            />
          </div>

          <div className="mb-3">
            <label className="block mb-1 font-medium" htmlFor="confirmPassword">
              Confirm Password:
            </label>
            <input
              type="password"
              id="confirmPassword"
              name="confirmPassword"
              value={formData.confirmPassword}
              onChange={handleChange}
              className="w-full p-2 border border-gray-300 rounded"
              required
            />
          </div>

          {/* Error Messages */}
          {messages.length > 0 && (
            <div className="mb-3">
              {messages.map((msg, index) => (
                <p key={index} className="text-red-500 text-sm">
                  {msg}
                </p>
              ))}
            </div>
          )}

          <div className="flex justify-center mb-3">
            <button
              type="submit"
              className="w-1/2 bg-green-600 text-white py-2 rounded hover:bg-green-500 transition"
            >
              Register
            </button>
          </div>

          <div className="text-center">
            Already have an account?{" "}
            <a href="/login" className="text-blue-600 hover:underline">
              Login
            </a>
          </div>
        </form>
      </div>
    </div>
  );
};

export default RegistrationForm;
