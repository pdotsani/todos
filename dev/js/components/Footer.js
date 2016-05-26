import React from 'react';

class Footer extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			website: "http://beego.me",
			email: "astaxie@gmail.com"
		};
	}

	render() {
		return (
			<div class="column small-12">
			  <footer class="row align-center text-center">
			    <div class="column medium-12 large-8">
			      Official website:
			      <a href={this.state.website}>{this.state.website}</a> /
			      Contact me:
			      <a href="mailto:{this.state.email}">{this.state.email}</a>
			    </div>
			  </footer>
			</div>
		)
	}
}

export default Footer;