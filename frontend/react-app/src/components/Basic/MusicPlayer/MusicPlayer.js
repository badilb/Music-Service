import React from 'react'
import "./musicplaya.scss"
import { FaPlay } from "react-icons/fa";


class MusicPlayer extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            showDescription: false, // Initially show the description
        };
        this.toggleDescription = this.toggleDescription.bind(this);
    }
    toggleDescription = () => {
        this.setState(prevState => ({ showDescription: !prevState.showDescription }));
    };

    render(){

        const { showDescription } = this.state.showDescription;
        return (
        <div className="player">
            <div className="head">
                <div className="back"></div>
                {/*<button onClick={() => this.setState({showDescription: !this.state.showDescription})}*/}
                {/*        className={"btn btn-danger"}>Toggle Description*/}
                {/*</button>*/}
                <button onClick={this.toggleDescription}
                        className={"btn btn-danger"}>Toggle Description
                </button>
                {this.state.showDescription && <div className="front">
                    <div className="avatar">
                        <img
                            src="http://vignette4.wikia.nocookie.net/lyricwiki/images/1/15/Semisonic_portrait.jpg/revision/latest?cb=20090409134209"
                            alt="Avatar"/>
                    </div>
                    <div className="infos">
                        <div className="titulo_song">
                            KFC - Killing Fabulous Chickens
                        </div>
                        <div className="duracao_song">
                            <i className="fa fa-clock-o"></i> Total time 45:12
                        </div>
                        <div className="tags">
                            <span>Educativo</span>
                            <span>Galinhas</span>
                            <span>Podcast</span>
                        </div>
                    </div>
                </div>}

            </div>
            <div className="timeline">
                <div className="soundline"></div>
                <div className="controllers active">
                    <div className="back"></div>
                    <div className="play"><FaPlay/></div>
                    <div className="forward"></div>
                </div>
            </div>
            <div className="rotation"></div>
        </div>


        )
    }


}

export default MusicPlayer