class Images {
	constructor() {
		this.imgUrl = "/api/images";
	}

	post(url) {
		new Promise(function(resolve, reject) {
			request.post(this.imgUrl)
				.send({url: url, alt: 'title'})
				.end(function(err, res) {
					if (err) reject(err);
					resolve();
				});
		});
	}

	getAll() {
		new Promise(function(resolve, reject) {
			request.get(this.imgUrl)
				.end(function(err, res) {
					if (err) reject(err);
					resolve(res);
				});
		});
	}
}

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

class Main extends React.Component {
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

ReactDOM.render(
	<Main />,
	document.getElementById('app')
);