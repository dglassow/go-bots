/*******************************************************************************\
|                                                                               |
| - Created by D.Glassow                                                        |
|                                                                               |
| - any variable or function inside will be available from any of your scripts. |
|                                                                               |
\*******************************************************************************/

// ###############################  Variables ###############################  // Comments are always here
minDeut = 25*1000*1000														   // Deut minimum for all bots.
me = (BotID)
myself = GetBotByID(me)
startTime = GetTimestamp()
myClass = CharacterClass()   
nextSleep = GetNextSleepTime()
now = GetTimestamp()
myHomeworld = GetHomeWorld()
bot = "bot"+myself.GetID()
statusID = myself.GetPlayerName()+".status"
botCoords = [
    "M:#:#:#",
]
sleepMinutes = sleepHours*60													// Global sleep minutes calculation
sleepSeconds = sleepMinutes*60													// Global sleep seconds calulation 
homeworld = GetHomeWorld()                                                     	// Where you are sending your resources from
slots = GetSlots()																// Get slots
now = GetTimeStamp()															/
// ###############################  Functions ###############################  //

// ################################  Dynamic ################################  //

/*******************************************************************************\
|                                                                               |
|                            End Script by D.Glassow                            |
|                                                                               |
\*******************************************************************************/
