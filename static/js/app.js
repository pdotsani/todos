"use strict";

var BeegoHeader = React.createClass({
	displayName: "BeegoHeader",

	render: function render() {
		return React.createElement(
			"div",
			{ "class": "column small-12" },
			React.createElement(
				"header",
				{ "class": "row align-center text-center" },
				React.createElement(
					"div",
					{ "class": "column medium-12 large-8" },
					React.createElement(
						"h1",
						null,
						"Welcome to Beego"
					),
					React.createElement(
						"h3",
						null,
						"yooo...."
					),
					React.createElement(
						"p",
						null,
						"Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra."
					)
				)
			)
		);
	}
});

var BeegoFooter = React.createClass({
	displayName: "BeegoFooter",

	getInitialState: function getInitialState() {
		return {
			website: "http://beego.me",
			email: "astaxie@gmail.com"
		};
	},

	render: function render() {
		return React.createElement(
			"div",
			{ "class": "column small-12" },
			React.createElement(
				"footer",
				{ "class": "row align-center text-center" },
				React.createElement(
					"div",
					{ "class": "column medium-12 large-8" },
					"Official website:",
					React.createElement(
						"a",
						{ href: this.state.website },
						this.state.website
					),
					" / Contact me:",
					React.createElement(
						"a",
						{ href: "mailto:{this.state.email}" },
						this.state.email
					)
				)
			)
		);
	}
});

var Main = React.createClass({
	displayName: "Main",

	render: function render() {
		return React.createElement(
			"div",
			null,
			React.createElement(BeegoHeader, null),
			React.createElement(BeegoFooter, null)
		);
	}
});

ReactDOM.render(React.createElement(Main, null), document.getElementById('app'));