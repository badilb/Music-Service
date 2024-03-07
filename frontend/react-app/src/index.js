import React from 'react';
import ReactDOM from 'react-dom/client';
import './css/index.css';
import HomePage from "./components/HomePage"
import {createBrowserRouter, RouterProvider, createRoutesFromElements, Route, Outlet} from "react-router-dom";
import ProfilePage from "./components/ProfilePage";
import NotFoundPage from "./components/Basic/NotFoundPage";
import Login from "./components/Auth/Login"
import Register from "./components/Auth/Register";
import Navbar from "./components/Basic/Navbar";
import {AuthProvider} from "./components/Auth/context/AuthProvider";
import Unauthorized from "./components/Basic/Unauthorized";
import RequireAuth from "./components/Basic/RequireAuth";
import Creator from "./components/Creator";
import MusicPlayer from "./components/Basic/MusicPlayer/MusicPlayer";
import MusicPlayer2 from "./components/Basic/MusicPlayer/MusicPlayer2";
import {MusicPlayerProvider} from "./components/Basic/MusicPlayer/MusicProvider/MusicProvider";

const root = ReactDOM.createRoot(document.getElementById('root'));


const router = createBrowserRouter(
    createRoutesFromElements(
        <Route element={<Layout />} errorElement={<NotFoundPage/>}>

            <Route path="/" element={<HomePage />} />
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login />} />
            <Route path="/unauthorized" element={<Unauthorized />} />
            <Route element={<RequireAuth allowedRoles={[1]} />}>
            {/*    zashishennoe   */}
                <Route path={"/creator"} element={<Creator />}> </Route>
            </Route>
            <Route element={<RequireAuth allowedRoles={[2]} />}>
                {/*    zashishennoe   */}
                <Route path={"/profile"} element={<ProfilePage />}> </Route>
            </Route>
        </Route>
    )
);
function Layout() {
    return (
        <div>
            <header>
                <Navbar />
            </header>
            <main>
                {/* 2️⃣ Render the app routes via the Layout Outlet */}

                <Outlet />
            </main>
            <footer> <MusicPlayer2 /> </footer>
        </div>
    );
}

root.render(
  <React.StrictMode>
      <MusicPlayerProvider>
          <AuthProvider>
              <RouterProvider router={router}/>
          </AuthProvider>
      </MusicPlayerProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
