import React from 'react';
import Images from '../services/Images';

class Board extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			images: []
		};
	}
	
	addImage() {
		var url = this.refs.inputurl.value;
		console.log("React: Add Image Executed!");
		console.log("image: ", url);
		Images
			.post(url)
			.then(function(data) {
				console.log("Request Success! ", data);
			})
			.catch(function(err) {
				console.error(err);
			});
	}

	render() {
		return (
			<div>
				<h2>Board</h2>
				<form onSubmit={this.addImage}>
					<input 
						type="text"
						placeholder="type image url here"
						ref='inputurl'/>
				</form>
				{this.images}
			</div>
		)
	}
}

export default Board;