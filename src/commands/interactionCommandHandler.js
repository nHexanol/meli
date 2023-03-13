const Discord = require('discord.js');
const stats = require('./stats.js');
const link = require('./link.js');
const parser = require('./parser.js');
const terrs = require('./territories.js');

module.exports = async function interactionCommandHandler(interaction) {
	switch (interaction.commandName) {
		case 'apply':
			// modals
			const applicationModal = new Discord.ModalBuilder()
				.setCustomId('application')
				.setTitle('Guild application')

			// username field
			const one = new Discord.TextInputBuilder()
				.setCustomId('one')
				.setLabel('In-game name and age')
				.setPlaceholder('Enter your username and age')
				.setStyle(Discord.TextInputStyle.Short)
				.setMinLength(5)
				.setMaxLength(20)
				.setRequired(true)

			// age field
			const two = new Discord.TextInputBuilder()
				.setCustomId('two')
				.setLabel('School and previous guild')
				.setPlaceholder('If you\'d like to share what you study, you can do that here!')
				.setStyle(Discord.TextInputStyle.Paragraph)
				.setRequired(true)

			// are you currently in school
			const three = new Discord.TextInputBuilder()
				.setCustomId('three')
				.setLabel('Guild wars')
				.setPlaceholder('Previous guild affiliation and are you willing to/learn to war and eco')
				.setStyle(Discord.TextInputStyle.Paragraph)
				.setRequired(true)

			const four = new Discord.TextInputBuilder()
				.setCustomId('four')
				.setLabel('What is your weekly playtime and classes')
				.setPlaceholder('Average playtime per week, what classes over level 100 do you have?')
				.setStyle(Discord.TextInputStyle.Paragraph)
				.setRequired(true)

			const five = new Discord.TextInputBuilder()
				.setCustomId('five')
				.setLabel('What are your favorite things to do on Wynn?')
				.setPlaceholder('Optional')
				.setRequired(false)
				.setStyle(Discord.TextInputStyle.Paragraph)


			const first = new Discord.ActionRowBuilder().addComponents(one);
			const second = new Discord.ActionRowBuilder().addComponents(two);
			const third = new Discord.ActionRowBuilder().addComponents(three);
			const fourth = new Discord.ActionRowBuilder().addComponents(four);
			const fifth = new Discord.ActionRowBuilder().addComponents(five);

			applicationModal.addComponents(first, second, third, fourth, fifth);

			await interaction.showModal(applicationModal);
			break;
		
		case 'stats':
			stats(interaction);
			break;
		case 'link':
			link(interaction);
			break;
		case 'terrs':
			var guild = interaction.options.getString('guild');
			terrs(guild, interaction);
			break;
		default:
			// manually parse the command
			parser(interaction);
			break;
	}

}