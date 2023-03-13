// importing all the shit
const { exec } = require('child_process');
const fs = require('fs');
const Discord = require('discord.js');
const client = new Discord.Client({ intents: Discord.IntentsBitField.Flags.Guilds | Discord.IntentsBitField.Flags.GuildMessages });
const interactionCommandHandler = require('./commands/interactionCommandHandler.js');
const modalExec = require('./events/modalSubmitExec.js');
const memberAdd = require('./events/joinMessage.js');
const processSelectMenu = require('./commands/statsSelectMenu.js');
const parser = require('./commands/parser.js');

// const values
const CLIENT_ID = '921868258502668318';
const TOKEN = 'OTIxODY4MjU4NTAyNjY4MzE4.GES1o6.zFg-_BvxeHHqbB0FGfzZV6eC8KsW2g9rWeSBx8'; // fs.readFileSync('./token.txt', 'utf8');

const commands = [
	new Discord.SlashCommandBuilder().setName('apply').setDescription('Start guild application forms'),
	new Discord.SlashCommandBuilder().setName('stats').setDescription('Query stats')
	.addStringOption(option => option.setName('query').setDescription('Username, UUID, guild or guild tag to query (optional)').setRequired(false)),
	new Discord.SlashCommandBuilder().setName('sp').setDescription('Soul point regen timer'),
	new Discord.SlashCommandBuilder().setName('terrs').setDescription('Show territories').addStringOption(option => option.setName('guild').setDescription('Guild to show').setRequired(false)).addStringOption(option => option.setName('area').setDescription('Crop to specific area').setRequired(false)),
	new Discord.SlashCommandBuilder().setName('guild').setDescription('Show guild stats').addStringOption(option => option.setName('guild').setDescription('Guild to show').setRequired(false)),
	new Discord.SlashCommandBuilder().setName('ping').setDescription('Ping!'),
	new Discord.SlashCommandBuilder().setName('info').setDescription('Get info about the bot'),
	new Discord.SlashCommandBuilder().setName('rebuild').setDescription('Force recompilation of the code and reload the bot').setDefaultMemberPermissions(Discord.PermissionFlagsBits.Administrator),
	new Discord.SlashCommandBuilder().setName('activity').setDescription('Show guild member activities (or your activities)')
	.addStringOption(option => option.setName('member').setDescription('Query specific guild member (username)').setRequired(false))
	.addUserOption(option => option.setName('user').setDescription('Query specific guild member (discord member)').setRequired(false)),
	new Discord.SlashCommandBuilder().setName('stop').setDescription('qwq').setDefaultMemberPermissions(Discord.PermissionFlagsBits.Administrator),
	new Discord.SlashCommandBuilder().setName('link').setDescription('Link discord account with username').setDefaultMemberPermissions(Discord.PermissionFlagsBits.PrioritySpeaker)
	.addStringOption(option => option.setName('username').setDescription('Username to link').setRequired(true).setMinLength(3).setMaxLength(16))
	.addUserOption(option => option.setName('discord').setDescription('Discord to link').setRequired(false)),

];
const rest = new Discord.REST({ version: '10' }).setToken(TOKEN);

// reload the cmds everytime so discord doesnt get ooga booga mad
(async () => {
	try {
		console.log('Started refreshing application (/) commands.');

		await rest.put(Discord.Routes.applicationCommands(CLIENT_ID), { body: commands });

		console.log('Successfully reloaded application (/) commands.');
	} catch (error) {
		console.error(error);
	}
})();

client.on('ready', function () {
	console.log('logged in');
	exec(`./events/logger.sh`, function (err, stdout, stderr) {
		if (err) {
			console.error(`exec error: ${err}`);
			return;
		}
	});
});

// listener for command and modals submit
client.on(Discord.Events.InteractionCreate, async function (interaction) {
	// if its / command then pass to the interactionCommandHandler
	if (interaction.isCommand()) await interactionCommandHandler(interaction);

	// if not / cmd then it's probably a modal or button event
	else if (interaction.isModalSubmit()) await processApplication(interaction);

	// if its not both then check if its a select menu
	else if (interaction.isStringSelectMenu()) await processSelectMenu(interaction);

	else if (interaction.isButton()) await parser(interaction);
});

async function processApplication(interaction) {

	if (interaction.customId === 'application') {
		// tell them
		modalExec(client, interaction);
	}
}

client.on(Discord.Events.GuildMemberAdd, async function (member) {
	member.roles.add('1058833931388076133');
	exec(`../bin/joinEmbed "${member.user.id}" "${member.user.avatarURL()}`, function (err, stdout, stderr) {
		if (err) {
			console.error(`exec error: ${err}`);
			return;
		}
	});
});

process.on('uncaughtException', function (err) {
	console.error(err);
});

client.login(TOKEN);