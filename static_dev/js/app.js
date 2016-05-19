var BeegoHeader = React.createClass({
	render: function() {
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
});

var BeegoFooter = React.createClass({
	getInitialState: function () {
		return {
			website: "http://beego.me",
			email: "astaxie@gmail.com"
		}
	},

	render: function () {
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
});

var Main = React.createClass({
	render: function() {
		return (
			<div>
				<BeegoHeader />
				<BeegoFooter />
			</div>	
		)
	}
});

ReactDOM.render(
	<Main />,
	document.getElementById('app')
);