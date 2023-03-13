module.exports = function (string) {
	// find username from string, it could be abc 20 or 20 abc
	const username = string.replace(',', '').split(' ').find(a => a.length > 2)
	return fetch(`https://api.wynncraft.com/v2/player/${username}/stats`)
		.then(res => res.json())
		.then(json => {
			return Promise.resolve(json.data[0].uuid);
	});

}