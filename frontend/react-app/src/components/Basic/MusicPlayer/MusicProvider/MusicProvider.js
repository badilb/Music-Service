import React, { createContext, useContext, useState } from 'react';

const MusicPlayerContext = createContext();

export const useMusicPlayer = () => {
    return useContext(MusicPlayerContext);
};

export const MusicPlayerProvider = ({ children }) => {
    const [audioUrl, setAudioUrl] = useState('');
    const [audioTitle, setAudioTitle] = useState('');

    const playAudio = (url, title) => {
        setAudioUrl(url);
        setAudioTitle(title);
    };

    return (
        <MusicPlayerContext.Provider value={{ audioUrl, audioTitle, playAudio }}>
            {children}
        </MusicPlayerContext.Provider>
    );
};