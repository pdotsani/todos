import React from 'react';

class Header extends React.Component {
	render() {
		return (
			<div class="column small-12">
			  <header class="row align-center text-center">
			    <div class="column medium-12 large-8">
			      <h1>Welcome to Beego</h1>
			      <h3>yooo....</h3>
			      <p>
			        Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
			      </p>
			    </div>
			  </header>
			</div>
		)
	}
}

export default Header;