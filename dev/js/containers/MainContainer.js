import React from 'react';
import Header from '../components/Header';
import Board from '../components/Board';
import Footer from '../components/Footer';

class MainContainer extends React.Component {
	render() {
		return (
			<div>
				<Header />
				<Board />
				<Footer />
			</div>	
		)
	}
}

export default MainContainer;