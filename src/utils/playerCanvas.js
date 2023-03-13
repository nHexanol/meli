const Discord = require('discord.js');
const { createCanvas, loadImage } = require('canvas');

module.exports = async function playerCanvas(p, interaction) {

	console.log('called playerCanvas.js');
	const canvas = createCanvas(1600, 900);
	const ctx = canvas.getContext('2d');

	await loadImage('../assets/0.png').then((image) => { ctx.drawImage(image, 0, 0, 1920, 1080); });
	var highestClass = Object.entries(p.data[0].characters)[0];
	console.log(highestClass[1])
	switch (highestClass[1].type) {
		case "MAGE":
			highestClassName = "Mage";
			break;
		case "SHAMAN":
			highestClassName = "Shaman";
			break;
		case "ASSASSIN":
			highestClassName = "Assassin";
			break;
		case "WARRIOR":
			highestClassName = "Warrior";
			break;
		case "ARCHER":
			highestClassName = "Archer";
			break;
		case "DARKWIZARD":
			highestClassName = "Dark Wizard";
			break;
		case "SKYSEER":
			highestClassName = "Skyseer";
			break;
		case "KNIGHT":
			highestClassName = "Knight";
			break;
		case "NINJA":
			highestClassName = "Ninja";
			break;
		case "HUNTER":
			highestClassName = "Hunter";
			break;
		default:
			highestClassName = "Unknown";
			break;
	}

	console.log(p.username);
	await loadImage(`https://mc-heads.net/avatar/${p.data[0].username}/256`)
		.then(function (image) {
			ctx.drawImage(image, canvas.width - 356, 75);
		});


	if (!p.data[0].meta.location.online) {
		var lastSeen = Date.now() - Date.parse(p.data[0].meta.lastJoin)
		years = Math.floor(lastSeen / (365 * 24 * 60 * 60 * 1000));
		days = Math.floor((lastSeen - years * (365 * 24 * 60 * 60 * 1000)) / (24 * 60 * 60 * 1000));
		hours = Math.floor((lastSeen - years * (365 * 24 * 60 * 60 * 1000) - days * (24 * 60 * 60 * 1000)) / (60 * 60 * 1000));
		minutes = Math.floor((lastSeen - years * (365 * 24 * 60 * 60 * 1000) - days * (24 * 60 * 60 * 1000) - hours * (60 * 60 * 1000)) / (60 * 1000));
		output = `${years > 0 ? years + " yr " : ""}${days > 0 ? days + " d " : ""}${hours > 0 ? hours + " hr " : ""}${minutes > 0 ? minutes + " min " : ""}`;
		output = (output[output.length - 1] == ":" ? output.slice(0, -1) : output).concat('ago');
	}
	else if (p.data[0].meta.location.online) output = "Online";

	ctx.font = 'bold 55pt Ubuntu';
	ctx.textAlign = 'left';
	ctx.fillStyle = '#ffffff';
	ctx.strokeStyle = '#000000';
	ctx.fillText(`${p.data[0].username}`, 75, 125);
	ctx.strokeText(`${p.data[0].username}`, 75, 125);
	var textWidth = ctx.measureText(p.data[0].username).width;
	ctx.font = '55pt Ubuntu';
	if (p.data[0].meta.tag.value != null) {
		if (p.data[0].meta.location.online) {
			ctx.fillText(`[${p.data[0].meta.tag.value}]  [${p.data[0].meta.location.server}]`, textWidth + 100, 125);
			ctx.strokeText(`[${p.data[0].meta.tag.value}]  [${p.data[0].meta.location.server}]`, textWidth + 100, 125);
		}
		else if (!p.data[0].meta.location.online) {
			ctx.fillText(`[${p.data[0].meta.tag.value}]`, textWidth + 100, 125);
			ctx.strokeText(`[${p.data[0].meta.tag.value}]`, textWidth + 100, 125);
		}
	}
	//guild and their rank
	if (p.data[0].guild.name == null) {
		var guild = "No guild"
	}
	else if (p.data[0].guild.name != null) {
		var guild = `${p.data[0].guild.rank}  of  ${p.data[0].guild.name}`
	}
	ctx.font = '35pt Ubuntu';
	ctx.fillText(`${guild}`, 75, 195);
	ctx.strokeText(`${guild}`, 75, 195);

	// informations
	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Total level  :', 75, 300);
	var textWidth = ctx.measureText("Total level  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`C ${p.data[0].global.totalLevel.combat.toLocaleString('en-US')} + P ${p.data[0].global.totalLevel.profession.toLocaleString('en-US')} = ${p.data[0].global.totalLevel.combined.toLocaleString('en-US')}`, 100 + textWidth, 300);
	ctx.strokeText(`C ${p.data[0].global.totalLevel.combat.toLocaleString('en-US')} + P ${p.data[0].global.totalLevel.profession.toLocaleString('en-US')} = ${p.data[0].global.totalLevel.combined.toLocaleString('en-US')}`, 100 + textWidth, 300);

	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Total playtime  :', 75, 375);
	var textWidth = ctx.measureText("Total playtime  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`${Math.round(Math.floor(p.data[0].meta.playtime) / 60 * 4.7).toLocaleString('en-US')} hours`, 100 + textWidth, 375);
	ctx.strokeText(`${Math.round(Math.floor(p.data[0].meta.playtime) / 60 * 4.7).toLocaleString('en-US')} hours`, 100 + textWidth, 375);

	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Total mobs killed  :', 75, 450);
	ctx.strokeText('Total mobs killed  :', 75, 450);
	var textWidth = ctx.measureText("Total mobs killed  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`${p.data[0].global.mobsKilled.toLocaleString('en-US')} mobs`, 100 + textWidth, 450);
	ctx.strokeText(`${p.data[0].global.mobsKilled.toLocaleString('en-US')} mobs`, 100 + textWidth, 450);

	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Not implemented  :', 75, 525);
	ctx.strokeText('Not implemented  :', 75, 525);
	var textWidth = ctx.measureText("Not implemented  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`<nil>`, 100 + textWidth, 525);
	ctx.strokeText(`<nil>`, 100 + textWidth, 525);

	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Logins/Deaths  :', 75, 600);
	ctx.strokeText('Logins/Deaths  :', 75, 600);
	var textWidth = ctx.measureText("Logins/Deaths  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`${p.data[0].global.logins.toLocaleString('en-US')}/${p.data[0].global.deaths.toLocaleString('en-US')} times`, 100 + textWidth, 600);
	ctx.strokeText(`${p.data[0].global.logins.toLocaleString('en-US')}/${p.data[0].global.deaths.toLocaleString('en-US')} times`, 100 + textWidth, 600);

	// last seen
	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Last seen  :', 75, 675);
	ctx.strokeText('Last seen  :', 75, 675);
	var textWidth = ctx.measureText("Last seen  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`${output}`, 100 + textWidth, 675);
	ctx.strokeText(`${output}`, 100 + textWidth, 675);

	var firstJoin = Date.now() - Date.parse(p.data[0].meta.firstJoin)
	years = Math.floor(firstJoin / (365 * 24 * 60 * 60 * 1000));
	days = Math.floor((firstJoin - years * (365 * 24 * 60 * 60 * 1000)) / (24 * 60 * 60 * 1000));
	hours = Math.floor((firstJoin - years * (365 * 24 * 60 * 60 * 1000) - days * (24 * 60 * 60 * 1000)) / (60 * 60 * 1000));
	minutes = Math.floor((firstJoin - years * (365 * 24 * 60 * 60 * 1000) - days * (24 * 60 * 60 * 1000) - hours * (60 * 60 * 1000)) / (60 * 1000));
	output = `${years > 0 ? years + " yr " : ""}${days > 0 ? days + " d " : ""}${hours > 0 ? hours + " hr " : ""}${minutes > 0 ? minutes + " min " : ""}`;
	output = output[output.length - 1] == ":" ? output.slice(0, -1) : output;

	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Joined  :', 75, 750);
	ctx.strokeText('Joined  :', 75, 750);
	var textWidth = ctx.measureText("Joined  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`${output}ago`, 100 + textWidth, 750);
	ctx.strokeText(`${output}ago`, 100 + textWidth, 750);

	ctx.font = 'bold 45pt Ubuntu';
	ctx.fillText('Highest leveled class  :', 75, 825);
	ctx.strokeText('Highest leveled class  :', 75, 825);
	var textWidth = ctx.measureText("Highest leveled class  :").width;
	ctx.font = '45pt Ubuntu';
	ctx.fillText(`${highestClassName} [${highestClass[1].professions.combat.level}/${highestClass[1].level.toLocaleString('en-US')}]`, 100 + textWidth, 825);
	ctx.strokeText(`${highestClassName} [${highestClass[1].professions.combat.level}/${highestClass[1].level.toLocaleString('en-US')}]`, 100 + textWidth, 825);

	ctx.font = '25pt Ubuntu';
	ctx.textAlign = 'right'
	ctx.fillStyle = "#fff";
	ctx.globalAlpha = 0.1;
	ctx.fillText(`nHexanol`, 1600, 25);
	ctx.globalAlpha = 1;

	
	const row = new Discord.ActionRowBuilder()
	/*
	if (p.data[0].guild.name !== null) {
		row.addComponents(
			new Discord.ButtonBuilder()
				.setLabel(`Query guild '${p.data[0].guild.name}'`)
				.setCustomId(`guild`)
				.setStyle(Discord.ButtonStyle.Primary)
		)
	}
	*/

	const option = [];
	console.log(p.data[0].characters)
	const chars = Object.entries(p.data[0].characters);
	for (var i in chars) {
		console.log();
		option.push({
			label: `${chars[i][1].type} [ Lv. ${chars[i][1].professions.combat.level} ]`,
			value: `${chars[i][0]}`,
			description: `Character UUID : ${chars[i][0]}`,
		});
	}
	row.addComponents(
		new Discord.StringSelectMenuBuilder()
			.setCustomId(`characters+${p.data[0].username}+${chars[i][0]}`)
			.setPlaceholder('Select a class to show more info')
			.addOptions(option)
	)

	console.log(interaction);
	await interaction.deferReply();
	await interaction.editReply({ files: [ canvas.toBuffer() ], components: [ row ] });
	setTimeout(() => {
		interaction.editReply({ files: [canvas.toBuffer()], components: [] });
	}, 60000);
}