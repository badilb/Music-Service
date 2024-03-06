import React from 'react'


class MusicPlayer extends React.Component{
    render(){return (
        <div>
        <div className="hero__backgorund owl-carousel">
            <div className="item" style="background: #3F51B5"></div>
            <div className="item" style="background: #8BC34A"></div>
            <div className="item" style="background: #673AB7"></div>
            <div className="item" style="background: #E91E63"></div>
        </div>



        <div className="main__container">
            <div className="contain__wrapper">
                <div className="music__wrapper">
                    <div className="left__panel">
                        <div className="cover owl-carousel">
                            <div className="item">
                                <img src="http://cdn.pikoff.com/wp-content/uploads/2016/05/02/Ai-Pink-and-purple-abstract-background-vector-free-download-350x350.jpg" alt=""/>
                            </div>
                            <div className="item">
                                <img src="http://cdn.pikoff.com/wp-content/uploads/2016/05/02/Ai-Pink-and-orange-abstract-background-vector-free-download-350x350.jpg" alt=""/>
                            </div>
                            <div className="item">
                                <img src="http://cdn.pikoff.com/wp-content/uploads/2016/05/02/Ai-Blue-abstract-background-vector-free-download-350x350.jpg" alt=""/>
                            </div>
                            <div className="item">
                                <img src="http://cdn.pikoff.com/wp-content/uploads/2016/05/02/Ai-Yellow-and-green-abstract-background-vector-free-download-350x350.jpg" alt=""/>
                            </div>
                        </div>
                    </div>

                    <div className="right__panel">
                        <div className="music__info__wrapper">
                            <div className="header">
                                <div className="icon__wrapper">

                                    <span className="icon-right"><i className="zmdi zmdi-apps"></i></span>
                                    <div className="title owl-carousel">
                                        <div className="item">
                                            <div className="song__name">Sorry The Movement</div>
                                            <p className="album__name">Justin Bieber</p>
                                        </div>
                                        <div className="item">
                                            <div className="song__name">Passenger Let Her Go</div>
                                            <p className="album__name">Passenger</p>
                                        </div>
                                        <div className="item">
                                            <div className="song__name">Counting Stars...</div>
                                            <p className="album__name">OneRepublic</p>
                                        </div>
                                        <div className="item">
                                            <div className="song__name">Waiting For Love</div>
                                            <p className="album__name">Avicii</p>
                                        </div>
                                    </div>

                                </div>
                            </div>

                            <div className="music__control">
                                <div className="music__button">
                                    <span className="button__prev"><i className="zmdi zmdi-fast-rewind"></i></span>
                                    <span className="button__pause"><i className="zmdi zmdi-pause"></i></span>
                                    <span className="button__next"><i className="zmdi zmdi-fast-forward"></i></span>
                                    <div>
                                        <input type="range" min="0" max="50" value="10" step="1" />
                                    </div>

                                </div>
                            </div>
                        </div>
                        <div className="music__menu">
                            <span className="icon-right"><i className="zmdi zmdi-close"></i></span>
                            <span className="menu_list"><i className="zmdi zmdi-favorite"></i>Faverate</span>
                            <span className="menu_list"><i className="zmdi zmdi-audio"></i>Music</span>
                            <span className="menu_list"><i className="zmdi zmdi-view-carousel"></i>Album</span>
                            <span className="menu_list"><i className="zmdi zmdi-time-restore"></i>History</span>
                            <span className="menu_list"><i className="zmdi zmdi-shopping-cart"></i>Cart</span>
                            <span className="menu_list"><i className="zmdi zmdi-settings"></i>Account</span>
                        </div>

                    </div>
                </div>
            </div>
        </div>
        </div>

    )
        }


}

export default MusicPlayer