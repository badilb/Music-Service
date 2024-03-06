import React, { useState, useEffect } from 'react';
import { FaPlay } from 'react-icons/fa';
import './musicplaya.scss'; // Make sure to import your styles

const MusicPlayer2 = () => {
    const [audioUrl, setAudioUrl] = useState('');
    const [audioTitle, setAudioTitle] = useState('');

    useEffect(() => {
        // Fetch audio data from your API as a Blob
        const fetchData = async () => {
            try {
                const response = await fetch('YOUR_AUDIO_API_URL'); // Replace with your API endpoint
                const audioBlob = await response.blob();
                setAudioUrl(URL.createObjectURL(audioBlob));

                // Assuming your API response has a title field
                const data = await response.json();
                setAudioTitle(data.title);
            } catch (error) {
                console.error('Error fetching audio data:', error);
            }
        };

        fetchData();
    }, []); // Empty dependency array ensures the effect runs only once on component mount

    const handlePlay = () => {
        // Add your logic to handle audio playback
        const audioElement = new Audio(audioUrl);
        audioElement.play();
    };

    return (
        <div className="player">
            <div className="title">{audioTitle}</div>
            <div className="timeline">
                <div className="soundline"></div>
                <div className="controllers active">
                    <div className="back"></div>
                    <div className="play" onClick={handlePlay}>
                        <FaPlay />
                    </div>
                    <div className="forward"></div>
                </div>
            </div>
        </div>
    );
};

export default MusicPlayer2;