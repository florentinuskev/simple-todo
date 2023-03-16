import { useState } from 'react'

// Routes
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import AuthRoute from './components/AuthRoute/AuthRoute';

// Pages
import LoginPage from './pages/LoginPage';
import MainPage from './pages/MainPage';

const router = createBrowserRouter([
  {
    path: "/",
    element: <AuthRoute>
      <MainPage />
    </AuthRoute>,
  },
  {
    path: "/login",
    element: <LoginPage />
  }
]);

function App() {

  return (
    <RouterProvider router={router} />
  )
}

export default App
