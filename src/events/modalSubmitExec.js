require('discord.js');
const { exec } = require('child_process');
const extract = require('../utils/extract.js');

module.exports = function sendData(client, interaction) {

	// extract data
	const userage = interaction.fields.getTextInputValue('one');
	const schoolprevguild = interaction.fields.getTextInputValue('two');
	const wars = interaction.fields.getTextInputValue('three');
	const ptclasses = interaction.fields.getTextInputValue('four');
	var fav = interaction.fields.getTextInputValue('five');

	if (fav.length == 0) {
		fav = '<nil>';
	}

    // since discord.js somehow cant build embed, we will do it in golang instead
	exec(`../bin/applicationEmbed "${interaction.user.id}" "${userage}" "${schoolprevguild}" "${wars}" "${ptclasses}" "${fav}" "${interaction.user.avatarURL()}"`, function (error, stdout, stderr) {
		if (error) {
			console.error(`exec error: ${error}`);
			interaction.reply('Application failed to submit!\n\nError while trying to execute `embed.go`\n```bash\n${error}\n```', { ephemeral: false });
			return;
		}
        else {
          interaction.reply('Application submitted!', { ephemeral: false });
					var uuid = '';
					extract(userage).then(function (data) { 
						uuid = data; 
						exec(`../bin/cacheUserData ${uuid} ${discordId}`, function (error, stdout, stderr) {
							if (error) {
								console.error(`exec error: ${error}`);
								return;
							}
						});	
					});
					var discordId = interaction.user.id;
					console.log(uuid);


        }
		console.log(`stdout: ${stdout}`);
		console.log(`stderr: ${stderr}`);
	});
}