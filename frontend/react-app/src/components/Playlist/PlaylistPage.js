// PlaylistPage.js
import React, { useState, useEffect } from 'react';
import axios from '../Auth/api/axios';
import { useParams } from 'react-router-dom'; // Import useParams for getting URL parameters

const PlaylistURL = "/playlists/${playlistId}/tracks"
const PlaylistPage = () => {
    const { playlistId } = useParams(); // Get the playlistId from the URL
    const [tracks, setTracks] = useState([]);

    useEffect(() => {
        // Replace 'YOUR_API_URL' with the actual URL of your API
        axios.get(PlaylistURL)
            .then(response => {
                setTracks(response.data); // Assuming your API returns an array of tracks
            })
            .catch(error => {
                console.error('Error fetching tracks:', error);
            });
    }, [playlistId]);

    const handlePlayAll = () => {
        // Add your logic to handle playing all tracks
        console.log('Play all tracks');
    };

    return (
        <div className="container mt-4">
            <h2>Playlist</h2>
            <h3>Tracks</h3>
            <button className="btn btn-primary" onClick={handlePlayAll}>Play All</button>
            <ul>
                {tracks.map(track => (
                    <li key={track.id}>{track.name}</li>
                ))}
            </ul>
        </div>
    );
};

export default PlaylistPage;