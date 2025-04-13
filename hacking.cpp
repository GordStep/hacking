#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <ostream>
#include <string>
#include <vector>
#include <chrono>
#include <thread>

using namespace std;


int
main ()
{
  vector<string> text = {
    " The BIOS initializes the hardware ",
    " A POST (Power-On Self Test) is performed ",
    " Defining processor initialization ",
    " Main Memory Identification (RAM) ",
    " Checking as hard drives and SSDs ",
    " Detecting and checking the keyboard ",
    " checking the operation of the system power supply ",
    " The processor goes into real mode ",
    " Loading SATA and IDE devices ",
    " Checking USB ports and other external devices ",
    " Initializing and checking the video card and monitor ",
    " Searching and reading the first sector of the boot device ",
    " Transferring control to the operating system loader ",
    " Reading the Boot Record ",
    " Loading the operating system kernel ",
    " Launching device drivers ",
    " Switching to protected processor mode ",
    " Initialization and configuration of ACPI ",
    " Downloading the file manager ",
    " Network connection, if configured ",
    " Displaying BIOS messages ",
    " Identification of removable media devices and their initialization ",
    " Performing BIOS updates, if available ",
    " Providing control to the operating system ",
    " Running background services and processes ",
    " Creating temporary files and cache ",
    " Hardware and software monitoring ",
    " Initialization of the graphics subsystem (if available) ",
    " Unpacking and preparing system files for operation ",
    " Drivers are updated if necessary ",
    " Preparing for the operation of the security system ",
    " Initializing and configuring the system sound ",
    " Activation of protective mechanisms ",
    " Notifying system administrators if configured ",
    " Downloading the authentication and authorization service ",
    " Activation of antivirus and antispyware programs ",
    " Loading local file resources ",
    " Activating Task Scheduler Services ",
    " Running background processes related to the network ",
    " Initialization of diagnostic and performance monitoring tools",
    " Disclosure of configuration information, if necessary ",
    " Checking the availability and status of intrusion protection devices ",
    " Initialization of temperature and performance monitoring ",
    " Loading the window manager ",
    " Checking the feedback from the equipment ",
    " Activation of device logging and monitoring tools ",
    " Restoring interrupted tasks or sessions ",
    " Checking the integrity of system files ",
    " Checking and launching backup and recovery tools ",
    " Activation of the audit and log analysis system ",
    " Preparing the user interface ",
    " Checking and starting the update service ",
    " Initializing and checking for security updates ",
    " Resolving resource conflicts or network connections ",
    " Checking the integrity of file systems ",
    " Search and launch software updates ",
    " Initialization and verification of the software environment ",
    " Launching data services ",
    " Checking and correcting user settings ",
    " Preparing and analyzing event logs ",
    " Checking for viruses and compromising software ",
    " Synchronization of time and date with network servers ",
    " Creating backups of system settings ",
    " Setting and configuring language and locale settings ",
    " Activating the mail service ",
    " Initialization of resource management software solutions ",
    " Launching programs related to project management ",
    " Initialization of energy monitoring and management tools ",
    " Launching programs for working with documents and office applications ",
    " Preparing the software interface ",
    " Activation of tools for working with multimedia and graphic files ",
    " Launch of communication tools and messengers ",
    " Initialization of the software interface for working with the Internet ",
    " Checking the software interface for working with databases ",
    " Activating and configuring development tools ",
    " Launching and verifying online services and cloud resources ",
    " Enabling multi-user mode ",
    " Downloading and configuring a subscriber line ",
    " Waiting for user input ",
    " Checking the operation and availability of critical updates ",
    " Activation of system monitoring mechanisms ",
    " Launching an anti-advertising and tracking program ",
    " Downloading data and information management programs ",
    " Initialization and verification of programs for working with sound ",
    " Running cycles of programs for analyzing and running network streams ",
    " Preparation of certification centers and key management ",
    " Checking and analyzing the word processing service ",
    " Restoring system settings ",
    " Activation of mechanisms for working with touch interfaces ",
    " Downloading data from the previous session ",
    " Confirmation and verification of the current configuration ",
    " Launching software update management systems ",
    " Checking and analyzing the integrity of user interface elements ",
    " Activating incident monitoring and management ",
    " Downloading and checking hardware and drivers ",
    " Initializing hardware interfaces ",
    " Activation and revision of security policies ",
    " Checking synchronization with cloud and Internet resources ",
    " Waiting and preparing for user access ",
    " Getting the ip address",
    " Connecting to minecraft servers",
    " Ask for permission to connect",
    " Stealing your cheesecakes",
    " Thinking about the future",
    " Recall the principle of action",
    " Drinking tea",
    " Today is a good day for Capricorns",
    " Сlearing the RAM",
    " Mom said at the bottom of the can",
    " Brother help me, I'm stuck!",
    " It's white",
    " A secret operation",
    " The data has been deleted",
    " Decrypting text files",
    " Ask mom to help"
  };
  char per = 'y';
  
  cout << "\b";
  cout << "Are you ready to start hacking?(Y/n) ";
  cin >> per;
  if (per == 'n')
    {
      cout << "Ok..." << endl;
      return 0;
    }
  else if (per == 'Y' || per == 'y' || per == '\n')
  {

    cout << "Loading";

    for (int i = 0; i < 3; i++)
    {
      for (int j = 0; j < 3; j++)
      {
        cout << "." << flush;
        this_thread::sleep_for(chrono::milliseconds(300));
      }
      cout << '\b' << '\b' << '\b';
    }
    cout << endl;

    for (;;)
    {
      int t = rand () % text.size ();
      int col = rand () % 10;
      if (col == 1)
      {
        cout << "[\033[37m  Load resurs  \033[0m] " << text[t] << endl;
        std::this_thread::sleep_for(std::chrono::milliseconds(500));
      }
      else if (col == 4)
      {
        cout << "[\033[31m   Stopping    \033[0m] " << text[t] << endl;
        std::this_thread::sleep_for(std::chrono::milliseconds(700));
      }
      else if (col == 6)
      {
        cout << "[\033[33m  Progressing  \033[0m] " << text[t] << endl;

        std::this_thread::sleep_for(std::chrono::milliseconds(600));
      }
      else
      {
        cout << "[\033[32m   Starting    \033[0m] " << text[t] << endl;
        std::this_thread::sleep_for(std::chrono::milliseconds(200));
      }
    }

    cout << endl;
  }

  else 
  {
    cout << "Invalid input." << endl;
  }

  return 0;
}
