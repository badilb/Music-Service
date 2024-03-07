import React , { useState, useEffect } from 'react'
import useAuth from "./Auth/hooks/useAuth";
import PlaylistComponent from "./Playlist/PlaylistComponent";

const HomePage = () => {
    const { auth } = useAuth();
    console.log(auth);

    return(
        <div>

            <h1> HOME PAGE </h1>
            <PlaylistComponent/>
        </div>
    )

}

export default HomePage;