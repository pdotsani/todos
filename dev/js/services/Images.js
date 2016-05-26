import * as request	 from 'superagent';

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

export default Images;