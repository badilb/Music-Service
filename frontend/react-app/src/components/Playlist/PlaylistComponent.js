// PlaylistComponent.js
import React, { useState, useEffect } from 'react';
import axios from '../Auth/api/axios';
import {Link} from 'react-router-dom'

const PlaylistsURL = "/api/playlists"
const PlaylistComponent = () => {
    const [playlists, setPlaylists] = useState([]);


    useEffect(() => {
        // Replace 'YOUR_API_URL' with the actual URL of your API
        axios.get(PlaylistsURL)
            .then(response => {
                setPlaylists(response); // Assuming your API returns an array of playlists
            })
            .catch(error => {
                console.error('Error fetching playlists:', error);
            });
    }, []); // The empty dependency array ensures the effect runs only once on component mount

    return (
        <div className="container mt-4">
            <h2>Playlists</h2>
            <div className="row">
                {playlists.map(playlist => (
                    <div key={playlist.ID} className="col-md-4 mb-4">
                        <div className="card">
                            {/* Add more card styling based on your needs */}
                            {/*<img src={playlislist.image} className="card-img-top" alt={playlist.Title} />*/}
                            <div className="card-body">
                                <h5 className="card-title">{playlist.Title}</h5>
                                {/*<p className="card-text">{playlist.description}</p>*/}
                                <Link to={`/playlist/${playlist.ID}`} className="btn btn-primary">View Playlist</Link>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default PlaylistComponent;