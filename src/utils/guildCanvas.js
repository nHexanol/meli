const Discord = require('discord.js');
const { createCanvas, loadImage } = require('canvas');

module.exports = async function guildCanvas(g, interaction) {
	console.log('called')
	if (!g.name) {
		interaction.reply(`Failed to query.`);
	}
	else {
		await interaction.deferReply();
		const canvas = createCanvas(1600, 900);
		const ctx = canvas.getContext('2d');

		// terr count is broken so we will send another api call
		var allTerrs = await fetch(`https://api.wynncraft.com/public_api.php?action=territoryList`).then(r => r.text())
		var territories = allTerrs.split(g.name).length - 1;

		await loadImage('../assets/0.png').then((image) => { ctx.drawImage(image, 0, 0, 1920, 1080); });

		//put all the canvas here
		ctx.font = 'bold 55pt Ubuntu';
		ctx.textAlign = 'left';
		ctx.fillStyle = '#fff';
		ctx.fillText(`${g.name}`, 75, 125);
		var textWidth = ctx.measureText(g.name).width;
		ctx.font = '55pt Ubuntu';
		ctx.fillText(`[${g.prefix}]`, 125 + textWidth, 125);
		ctx.font = '35pt Ubuntu';
		ctx.fillText(`Owned by `, 75, 185);
		var textWidth = ctx.measureText('Owned by ').width;
		ctx.font = 'bold 35pt Ubuntu';
		ctx.fillText(`${g.members.find(m => m.rank === 'OWNER').name}`, 75 + textWidth, 185);

		ctx.font = 'bold 45pt Ubuntu';
		var textWidth = ctx.measureText("Level  :").width;
		ctx.fillText('Level  :', 75, 300);
		ctx.font = '45pt Ubuntu';
		ctx.fillText(`${g.level}  (${g.xp * 10}%)`, 125 + textWidth, 300);

		ctx.font = 'bold 45pt Ubuntu';
		var textWidth = ctx.measureText("Created  :").width;
		ctx.fillText('Created  :', 75, 375);
		ctx.font = '45pt Ubuntu';
		ctx.fillText(`${g.createdFriendly}`, 125 + textWidth, 375);

		ctx.font = 'bold 45pt Ubuntu';
		var textWidth = ctx.measureText("Total members  :").width;
		ctx.fillText('Total members  :', 75, 450);
		ctx.font = '45pt Ubuntu';
		ctx.fillText(`${g.members.length}`, 125 + textWidth, 450);

		ctx.font = 'bold 45pt Ubuntu';
		var textWidth = ctx.measureText("Total territories  :").width;
		ctx.fillText('Total territories  :', 75, 525);
		ctx.font = '45pt Ubuntu';
		ctx.fillText(`${territories}`, 125 + textWidth, 525);


		//credits
		ctx.font = '25pt Ubuntu';
		ctx.textAlign = 'right'
		ctx.fillStyle = "#fff";
		ctx.globalAlpha = 0.25;
		ctx.fillText(`nHexanol`, 1600, 25);
		ctx.globalAlpha = 1;

		const row = new Discord.ActionRowBuilder()
			.addComponents(
				new Discord.ButtonBuilder()
					.setCustomId(`onlinemember_${g.name.replace(/ /g, '-')}`)
					.setLabel('Query online members')
					.setStyle(Discord.ButtonStyle.Primary),
			)
			.addComponents(
				new Discord.ButtonBuilder()
					.setCustomId(`xp_${g.name.replace(/ /g, '-')}`)
					.setLabel('XP leaderboard')
					.setStyle(Discord.ButtonStyle.Primary),
			)
			.addComponents(
				new Discord.ButtonBuilder()
					.setCustomId(`activities_${g.name.replace(/ /g, '-')}`)
					.setLabel('Member Activities')
					.setStyle(Discord.ButtonStyle.Primary),
			)

		if (territories > 0) {
			row.addComponents(
				new Discord.ButtonBuilder()
					.setCustomId(`territories_${g.name.replace(/ /g, '-')}`)
					.setLabel('Show territories')
					.setStyle(Discord.ButtonStyle.Secondary)
			)
		}

		await interaction.editReply({ files: [canvas.toBuffer()], components: [ row ] });
		setTimeout(() => {
			// remove buttons
			interaction.editReply({ files: [canvas.toBuffer()], components: [] });
		},60000)
	}
}