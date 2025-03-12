import React from "react";

const LandingPage: React.FC = () => {
  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-200 text-center">
      <div className="bg-white p-10 rounded-lg shadow-md">
        <h1 className="text-3xl font-bold mb-5">Welcome to ToDo App</h1>
        <p className="mb-7 text-gray-700">
          Organize your tasks efficiently and boost your productivity. Sign up
          to get started or log in to access your tasks.
        </p>
        <div className="space-x-4">
          <a
            href="/register"
            className="bg-green-600 text-white px-5 py-2 rounded-md hover:bg-green-700 transition"
          >
            Sign Up
          </a>
          <a
            href="/login"
            className="bg-blue-600 text-white px-5 py-2 rounded-md hover:bg-blue-700 transition"
          >
            Log In
          </a>
        </div>
      </div>
    </div>
  );
};

export default LandingPage;
