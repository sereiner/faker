package company

var formats = []string{
	`${companyPrefix}${companySuffix}`,
}

var companySuffix = []string{"科技", "网络", "信息", "传媒"}

var companyPrefix = []string{
	"超艺", "和泰", "九方", "鑫博腾飞", "戴硕电子",
	"济南亿次元", "海创", "创联世纪", "凌云", "泰麒麟",
	"彩虹", "兰金电子", "晖来计算机", "天益", "恒聪百汇",
	"菊风公司", "惠派国际公司", "创汇", "思优", "时空盒数字",
	"易动力", "飞海科技", "华泰通安", "盟新", "商软冠联",
	"图龙信息", "易动力", "华远软件", "创亿", "时刻",
	"开发区世创", "明腾", "良诺", "天开", "毕博诚", "快讯",
	"凌颖信息", "黄石金承", "恩悌", "雨林木风计算机",
	"双敏电子", "维旺明", "网新恒天", "数字100", "飞利信",
	"立信电子", "联通时科", "中建创业", "新格林耐特",
	"新宇龙信息", "浙大万朋", "MBP软件", "昂歌信息",
	"万迅电脑", "方正科技", "联软", "七喜", "南康", "银嘉",
	"巨奥", "佳禾", "国讯", "信诚致远", "浦华众城", "迪摩",
	"太极", "群英", "合联电子", "同兴万点", "襄樊地球村",
	"精芯", "艾提科信", "昊嘉", "鸿睿思博", "四通", "富罳",
	"商软冠联", "诺依曼软件", "东方峻景", "华成育卓", "趋势",
	"维涛", "通际名联",
}

var catchPhrase = []string{
	"燕舞，燕舞，一曲歌来一片情。",
	"康师傅方便面，好吃看得见。",
	"不要太潇洒！",
	"让一亿人先聪明起来。",
	"共创美的前程，共度美的人生。",
	"省优，部优，葛优？",
	"喝孔府宴酒，做天下文章。",
	"健康成就未来。",
	"牙好，胃口就好，身体倍儿棒，吃嘛嘛香。",
	"永远的绿色，永远的秦池。",
	"坐红旗车，走中国路。",
	"要想皮肤好，早晚用大宝。",
	"孔府家酒，叫人想家。",
	"补钙新观念，吸收是要害。",
	"喝汇源果汁，走健康之路。",
	"爱的就是你!",
	"一种可以世袭的古典浪漫",
	"实力创造价值",
	"爱生活，爱拉芳！",
	"人类失去联想，世界将会怎样？",
	"做女人挺好！",
	"世界在你眼中？",
	"今天你有否亿唐？",
	"只溶在口，不溶在手。",
	"三千烦恼丝，健康新开始。",
	"维维豆奶，欢乐开怀。",
	"我们的光彩来自你的风采。",
	"钻石恒久远，一颗永流传。",
	"放我的真心在你的手心。",
	"小身材，大味道。",
	"牛奶香浓，丝般感受。",
	"聆听并不代表沉默，有时安静也是一种力量。",
	"滴滴香浓，意犹未尽。",
	"水晶之恋，一生不变。",
	"中国移动通信，沟通从心开始！",
	"网易，网聚人的力量！",
	"科技以人为本，诺基亚",
	"我们一直在努力！",
	"阳光总在风雨后",
	"男人对西服的要求，就是女人对男人的要求",
	"晚报，不晚报",
	"原来生活可以更美的",
	"明天的明天，你还会送我“水晶之恋”吗？",
	"卫浴出出进进的快感",
	"有家就有联合利华",
	"减脂减肥，其实是一种生活态度",
	"人头马一开，好事自然来。",
	"假如五指一样长，怎能满足用户不同需求？",
	"新飞广告做的好，不好新飞冰箱好",
	"传奇品质，百年张裕",
	"李宁：把出色留给自己",
	"一旦拥有，别无选择",
	"科技让你更轻松",
	"情系中国结，联通四海心",
	"海尔，中国造",
	"SOHU：足迹生活每一天",
	"果冻我要喜之郎",
	"国宝大熊猫，心纯天自高",
	"世界因为不同",
	"放低偏见，你会有出色发现！",
	"Just",
	"创意似金，敬业如牛",
	"不要让男人一手把握",
	"如同情人的手",
	"金窝银窝，不如自己的安乐窝。",
	"没有什么大不了的",
	"时间因我存在",
	"只要有梦想",
	"南方周末",
	"时间改变一切",
	"地球人都知道了",
	"众里寻他千百度，想要几度就几度",
	"您身边的银行，可信赖的银行",
	"三叶钢琴：学琴的孩子不会变坏",
	"柯达：串起生活每一刻",
	"大众甲克虫汽车：想想还是小的好",
	"一直被模拟,从未被超越",
	"幸福生活",
	"朗讯的创造力科技的原动力",
	"事事因你而出色",
	"运动之美，世界共享",
	"鹤舞白沙",
	"想知道“清嘴”的味道吗？",
	"弹指一挥间，世界皆互联",
	"更多选择、更多欢笑",
	"方太，让家的感觉更好",
	"世上仅此一件，今生与你结缘！",
	"白里透红与众不同",
	"没有蛀牙-佳洁士",
	"有线的价值",
	"享受快乐科技",
	"四海一家的解决之道",
	"娃哈哈纯净水：爱你等于爱自己",
	"农民山泉：有点甜",
	"博大精深，西门子",
	"一切尽在把握",
	"声声百思特，遥遥两相知",
	"一呼天下应",
	"让我们做得更好！",
	"暖和亲情，金龙鱼的大家庭。",
	"自然最健康，绿色好心情",
	"支起网络世界",
	"立邦漆：处处放光彩！",
	"fm365:真情互动！",
	"庄重一生，吉祥一生。",
	"人人都为礼品愁，我送北极海狗油。",
	"假如说人生的离合是一场戏，那么百年的好合更是早有安排！",
	"一品黄山天高云淡",
	"上上下下的享受！",
	"我是、我行、我素",
	"让无力者有力，让悲观者前行",
	"金利来—-男人的世界！",
	"百衣百顺",
	"聪明何必绝顶，慧根长留",
	"水往高处流",
	"大石化小，小石化了！",
	"“闲”妻良母",
	"“口服”，“心服”！",
	"盛满青春的秘密！",
	"三十六计走为上",
	"为了她的节日，献上您纯金般的心！",
	"用我们的钓线，你可以在鱼儿发现你之前先找到它",
	"生活就是一场运动，喝下它。",
	"选择维聚阿尔，已经表明你心明眼亮。",
	"佳能，我们看得见你想表达什么。",
	"天天都是春天",
	"假如你不来，广告明星就是他",
	"享受黑夜中偷拍的快感！",
	"彩信发送动人一刻",
	"灵感点亮生活!",
	"聪明演绎，无处不在！",
	"事业我一定争取，对你我从未放弃!",
	"波导手机，手机中的战斗机",
	"鄂尔多斯羊绒衫暖和全世界",
	"洁婷245再大的动作也不要紧",
	"做光明的牛，产光明的奶",
	"假如你的汽车会游泳的话，请照直开，不必刹车。",
	"永远要让驾驶执照比你自己先到期。",
	"请记住，上帝并不是十全十美的，它给汽车预备了备件，而人没有。",
	"小别意酸酸，欢聚心甜甜。",
	"除钞票外，承印一切。",
	"更多欢乐，更多选择",
	"美由你做主",
	"由我天地宽",
	"Sun是太阳，Java是月亮。",
	"不断创新，因为专心",
	"趁早下『斑』，请勿『痘』留。",
	"请不要同刚刚走出本院的女人调情，她或许就是你的外祖母。",
	"创新就是生活",
	"有一个漂亮的地方，万科四季花城",
	"建筑无限生活",
	"臭名远扬，香飘万里",
	"尝尝欢笑，经常麦当劳",
	"深入成就深度",
	"出色湖南，红网了然！",
	"因为网络，地球如村！",
	"一种质感",
	"恒久期盼",
	"繁荣民族文化",
	"不信，死给你看！",
	"天生的，强生的",
	"雪津啤酒，真情的味道！",
	"听世界，打天下",
	"雅芳比女人更了解女人",
	"Sun是太阳，Java是月亮。",
	"中国网通",
	"无线你的无限",
	"家有三洋，冬暖夏凉",
	"倾诉冬日暖语",
	"谁让我心动？",
	"灵活，让篮球场不再是一个平面",
	"别吻我，我怕修。",
	"一呼四应！",
	"无所不包！",
	"当之无愧",
	"以帽取人！",
	"一毛不拔！",
	"自讨苦吃！",
	"成功与科技共辉映",
	"没有最",
}

var jobTitleFormat = []string{
	"Able Seamen", "Account Manager", "Accountant", "Actor", "Actuary", "Adjustment Clerk", "Admin", "Administrative Law Judge", "Administrative Services Manager", "Administrative Support Supervisors", "Advertising Manager OR Promotions Manager", "Advertising Sales Agent", "Aerospace Engineer", "Agricultural Crop Farm Manager", "Agricultural Crop Worker", "Agricultural Engineer", "Agricultural Equipment Operator", "Agricultural Inspector", "Agricultural Manager", "Agricultural Product Grader Sorter", "Agricultural Sales Representative", "Agricultural Science Technician", "Agricultural Sciences Teacher", "Agricultural Technician", "Agricultural Worker", "Air Crew Member", "Air Crew Officer", "Air Traffic Controller", "Aircraft Assembler", "Aircraft Body Repairer", "Aircraft Cargo Handling Supervisor", "Aircraft Engine Specialist", "Aircraft Launch and Recovery Officer", "Aircraft Launch Specialist", "Aircraft Mechanics OR Aircraft Service Technician", "Aircraft Rigging Assembler", "Aircraft Structure Assemblers", "Airfield Operations Specialist", "Airframe Mechanic", "Airline Pilot OR Copilot OR Flight Engineer", "Algorithm Developer", "Alteration Tailor", "Ambulance Driver", "Amusement Attendant", "Anesthesiologist", "Animal Breeder", "Animal Care Workers", "Animal Control Worker", "Animal Husbandry Worker", "Animal Scientist", "Animal Trainer", "Annealing Machine Operator", "Announcer", "Answering Service", "Anthropologist", "Anthropologist OR Archeologist", "Anthropology Teacher", "Appliance Repairer", "Arbitrator", "Archeologist", "Architect", "Architectural Drafter", "Architectural Drafter OR Civil Drafter", "Architecture Teacher", "Archivist", "Armored Assault Vehicle Crew Member", "Armored Assault Vehicle Officer", "Art Director", "Art Teacher", "Artillery Officer", "Artillery Crew Member", "Artist", "Assembler", "Assessor", "Astronomer", "Athletes and Sports Competitor", "Athletic Trainer", "Atmospheric and Space Scientist", "Audio and Video Equipment Technician", "Audiologist", "Audio-Visual Collections Specialist", "Auditor", "Auditor", "Automatic Teller Machine Servicer", "Automotive Body Repairer", "Automotive Glass Installers", "Automotive Master Mechanic", "Automotive Mechanic", "Automotive Specialty Technician", "Automotive Technician", "Auxiliary Equipment Operator", "Aviation Inspector", "Avionics Technician",
	"Bailiff", "Baker", "Barber", "Bartender", "Bartender Helper", "Battery Repairer", "Bellhop", "Bench Jeweler", "Benefits Specialist", "Bicycle Repairer", "Bill and Account Collector", "Bindery Machine Operator", "Bindery Worker", "Biochemist", "Biochemist or Biophysicist", "Biological Science Teacher", "Biological Scientist", "Biological Technician", "Biologist", "Biomedical Engineer", "Biophysicist", "Board Of Directors", "Boat Builder and Shipwright", "Boiler Operator", "Boilermaker", "Bookbinder", "Bookkeeper", "Brake Machine Setter", "Brattice Builder", "Brazer", "Brazing Machine Operator", "Brickmason", "Bridge Tender OR Lock Tender", "Broadcast News Analyst", "Broadcast Technician", "Brokerage Clerk", "Budget Analyst", "Buffing and Polishing Operator", "Building Cleaning Worker", "Building Inspector", "Bulldozer Operator", "Bus Driver", "Business Development Manager", "Business Manager", "Business Operations Specialist", "Business Teacher", "Butcher", "Buyer",
	"Cabinetmaker", "Cafeteria Cook", "Calibration Technician OR Instrumentation Technician", "Camera Operator", "Camera Repairer", "Captain", "Caption Writer", "Cardiovascular Technologist", "Career Counselor", "Carpenter", "Carpenter Assembler and Repairer", "Carpet Installer", "Cartographer", "Cartoonist", "Carver", "Cashier", "Casting Machine Operator", "Casting Machine Set-Up Operator", "ccc", "Ceiling Tile Installer", "Cement Mason and Concrete Finisher", "Central Office", "Central Office and PBX Installers", "Central Office Operator", "CEO", "CFO", "Chef", "Chemical Engineer", "Chemical Equipment Controller", "Chemical Equipment Operator", "Chemical Equipment Tender", "Chemical Plant Operator", "Chemical Technician", "Chemist", "Chemistry Teacher", "Child Care", "Child Care Worker", "Chiropractor", "Choreographer", "City", "City Planning Aide", "Civil Drafter", "Civil Engineer", "Civil Engineering Technician", "Claims Adjuster", "Claims Examiner", "Claims Taker", "Cleaners of Vehicles", "Clergy", "Clerk", "Clinical Laboratory Technician", "Clinical Psychologist", "Clinical School Psychologist", "Coaches and Scout", "Coating Machine Operator", "Coil Winders", "Command Control Center Officer", "Command Control Center Specialist", "Commercial and Industrial Designer", "Commercial Diver", "Commercial Pilot", "Communication Equipment Repairer", "Communication Equipment Worker", "Communications Equipment Operator", "Communications Teacher", "Community Service Manager", "Compacting Machine Operator", "Compensation and Benefits Manager", "Compliance Officers", "Composer", "Computer", "Computer Hardware Engineer", "Computer Operator", "Computer Programmer", "Computer Repairer", "Computer Science Teacher", "Computer Scientist", "Computer Security Specialist", "Computer Software Engineer", "Computer Specialist", "Computer Support Specialist", "Computer Systems Analyst", "Computer-Controlled Machine Tool Operator", "Concierge", "Conservation Scientist", "Construction", "Construction Carpenter", "Construction Driller", "Construction Equipment Operator", "Construction Laborer", "Construction Manager", "Continuous Mining Machine Operator", "Control Valve Installer", "Conveyor Operator", "Cook", "Cooling and Freezing Equipment Operator", "Copy Machine Operator", "Copy Writer", "Coremaking Machine Operator", "Coroner", "Corporate Trainer", "Correctional Officer", "Correspondence Clerk", "Cost Estimator", "Costume Attendant", "Counseling Psychologist", "Counselor", "Counsil", "Courier", "Court Clerk", "Court Reporter", "Craft Artist", "Crane and Tower Operator", "Creative Writer", "Credit Checkers Clerk", "Credit Analyst", "Credit Authorizer", "Credit Checker", "Criminal Investigator", "Crossing Guard", "Crushing Grinding Machine Operator", "CSI", "CTO", "Cultural Studies Teacher", "Curator", "Custom Tailor", "Customer Service Representative", "Cutting Machine Operator", "Cutting Machine Operator",
	"Dancer", "Data Entry Operator", "Data Processing Equipment Repairer", "Database Administrator", "Database Manager", "Deburring Machine Operator", "Decorator", "Dental Assistant", "Dental Hygienist", "Dental Laboratory Technician", "Dentist", "Designer", "Desktop Publisher", "Detective", "Diagnostic Medical Sonographer", "Diamond Worker", "Diesel Engine Specialist", "Dietetic Technician", "Director Of Business Development", "Director Of Marketing", "Director Of Social Media Marketing", "Director Of Talent Acquisition", "Director Religious Activities", "Directory Assistance Operator", "Dishwasher", "Dispatcher", "Distribution Manager", "Door To Door Sales", "Dot Etcher", "Drafter", "Dragline Operator", "Dredge Operator", "Drilling and Boring Machine Tool Setter", "Driver-Sales Worker", "Drycleaning Machine Operator", "Drywall Ceiling Tile Installer", "Drywall Installer",
	"Earth Driller", "Economics Teacher", "Economist", "Editor", "Education Administrator", "Education Teacher", "Educational Counselor OR Vocationall Counselor", "Educational Psychologist", "Electric Meter Installer", "Electric Motor Repairer", "Electrical and Electronic Inspector and Tester", "Electrical and Electronics Drafter", "Electrical Drafter", "Electrical Engineer", "Electrical Engineering Technician", "Electrical Parts Reconditioner", "Electrical Power-Line Installer", "Electrical Sales Representative", "Electrician", "Electrician", "Electrolytic Plating Machine Operator", "Electromechanical Equipment Assembler", "Electro-Mechanical Technician", "Electronic Drafter", "Electronic Engineering Technician", "Electronic Equipment Assembler", "Electronic Masking System Operator", "Electronics Engineer", "Electronics Engineering Technician", "Electrotyper", "Elementary and Secondary School Administrators", "Elementary School Teacher", "Elevator Installer and Repairer", "Eligibility Interviewer", "Embalmer", "Embossing Machine Operator", "Emergency Management Specialist", "Emergency Medical Technician and Paramedic", "Employment Interviewer", "Engine Assembler", "Engineer", "Engineering", "Engineering Manager", "Engineering Teacher", "Engineering Technician", "English Language Teacher", "Engraver", "Entertainer and Performer", "Entertainment Attendant", "Environmental Compliance Inspector", "Environmental Engineer", "Environmental Engineering Technician", "Environmental Science Teacher", "Environmental Science Technician", "Environmental Scientist", "Epidemiologist", "Equal Opportunity Representative", "Etcher", "Etcher and Engraver", "Event Planner", "Excavating Machine Operator", "Executive Secretary", "Exhibit Designer", "Explosives Expert", "Extraction Worker", "Extruding and Drawing Machine Operator", "Extruding Machine Operator",
	"Fabric Mender", "Fabric Pressers", "Farm and Home Management Advisor", "Farm Equipment Mechanic", "Farm Labor Contractor", "Farmer", "Farmworker", "Fashion Designer", "Fashion Model", "Fast Food Cook", "Fence Erector", "Fiber Product Cutting Machine Operator", "Fiberglass Laminator and Fabricator", "File Clerk", "Film Laboratory Technician", "Financial Analyst", "Financial Examiner", "Financial Manager", "Financial Services Sales Agent", "Financial Specialist", "Fire Fighter", "Fire Inspector", "Fire Investigator", "Fire-Prevention Engineer", "First-Line Supervisor-Manager of Landscaping, Lawn Service, and Groundskeeping Worker", "Fish Game Warden", "Fish Hatchery Manager", "Fishery Worker", "Fishing OR Forestry Supervisor", "Fitness Trainer", "Fitter", "Flight Attendant", "Floor Finisher", "Floor Layer", "Floral Designer", "Food Batchmaker", "Food Cooking Machine Operators", "Food Preparation", "Food Preparation and Serving Worker", "Food Preparation Worker", "Food Science Technician", "Food Scientists and Technologist", "Food Servers", "Food Service Manager", "Food Tobacco Roasting", "Foreign Language Teacher", "Forensic Investigator", "Forensic Science Technician", "Forest and Conservation Technician", "Forest and Conservation Worker", "Forest Fire Fighter", "Forest Fire Fighting Supervisor", "Forest Fire Inspector", "Forester", "Forestry Conservation Science Teacher", "Forging Machine Setter", "Forming Machine Operator", "Forming Machine Operator", "Foundry Mold and Coremaker", "Fraud Investigator", "Freight Agent", "Freight and Material Mover", "Freight Inspector", "Funeral Attendant", "Funeral Director", "Furnace Operator", "Furniture Finisher",
	"Gaming Cage Worker", "Gaming Dealer", "Gaming Manager", "Gaming Service Worker", "Gaming Supervisor", "Gaming Surveillance Officer", "Garment", "Gas Appliance Repairer", "Gas Compressor Operator", "Gas Distribution Plant Operator", "Gas Plant Operator", "Gas Processing Plant Operator", "Gas Pumping Station Operator", "Gas Pumping Station Operator", "Gauger", "GED Teacher", "General Farmworker", "General Manager", "General Practitioner", "Geographer", "Geography Teacher", "Geological Data Technician", "Geological Sample Test Technician", "Geologist", "Geoscientists", "Glass Blower", "Glass Cutting Machine Operator", "Glazier", "Gluing Machine Operator", "Government", "Government Property Inspector", "Government Service Executive", "Graduate Teaching Assistant", "Graphic Designer", "Grinder OR Polisher", "Grinding Machine Operator", "Grips", "Grounds Maintenance Worker",
	"Hairdresser OR Cosmetologist", "Hand Trimmer", "Hand Presser", "Hand Sewer", "Hazardous Materials Removal Worker", "Head Nurse", "Health Educator", "Health Practitioner", "Health Services Manager", "Health Specialties Teacher", "Health Technologist", "Healthcare", "Healthcare Practitioner", "Healthcare Support Worker", "Heat Treating Equipment Operator", "Heaters", "Heating and Air Conditioning Mechanic", "Heating Equipment Operator", "Heavy Equipment Mechanic", "Highway Maintenance Worker", "Highway Patrol Pilot", "Historian", "History Teacher", "Hoist and Winch Operator", "Home", "Home Appliance Installer", "Home Appliance Repairer", "Home Economics Teacher", "Home Entertainment Equipment Installer", "Home Health Aide", "Homeland Security", "Horticultural Worker", "Host and Hostess", "Hotel Desk Clerk", "House Cleaner", "Housekeeper", "Housekeeping Supervisor", "HR Manager", "HR Specialist", "Human Resource Director", "Human Resource Manager", "Human Resources Assistant", "Human Resources Manager", "Human Resources Specialist", "Hunter and Trapper", "HVAC Mechanic", "Hydrologist",
	"Illustrator", "Immigration Inspector OR Customs Inspector", "Industrial Engineer", "Industrial Engineering Technician", "Industrial Equipment Maintenance", "Industrial Machinery Mechanic", "Industrial Production Manager", "Industrial Safety Engineer", "Industrial-Organizational Psychologist", "Infantry", "Infantry Officer", "Information Systems Manager", "Inspector", "Installation and Repair Technician", "Instructional Coordinator", "Instrument Sales Representative", "Insulation Installer", "Insulation Worker", "Insurance Investigator", "Insurance Appraiser", "Insurance Claims Clerk", "Insurance Policy Processing Clerk", "Insurance Sales Agent", "Insurance Underwriter", "Interaction Designer", "Interior Designer", "Internist", "Interpreter OR Translator", "Interviewer", "Irradiated-Fuel Handler",
	"Janitor", "Janitorial Supervisor", "Jeweler", "Jewelry Model OR Mold Makers", "Job Printer", "Judge",
	"Keyboard Instrument Repairer and Tuner", "Kindergarten Teacher",
	"Landscape Architect", "Landscape Artist", "Landscaper", "Landscaping", "Lathe Operator", "Laundry OR Dry-Cleaning Worker", "Law Clerk", "Law Enforcement Teacher", "Law Teacher", "Lawn Service Manager", "Lawyer", "Lay-Out Worker", "Legal Secretary", "Legal Support Worker", "Legislator", "Letterpress Setters Operator", "Librarian", "Library Assistant", "Library Science Teacher", "Library Technician", "Library Worker", "License Clerk", "Licensed Practical Nurse", "Licensing Examiner and Inspector", "Life Science Technician", "Life Scientists", "Lifeguard", "Loading Machine Operator", "Loan Counselor", "Loan Interviewer", "Loan Officer", "Locker Room Attendant", "Locksmith", "Locomotive Engineer", "Locomotive Firer", "Lodging Manager", "Log Grader and Scaler", "Logging Equipment Operator", "Logging Supervisor", "Logging Tractor Operator", "Logging Worker", "Logistician",
	"Machine Feeder", "Machine Operator", "Machine Tool Operator", "Machinery Maintenance", "Machinist", "Maid", "Mail Clerk", "Mail Machine Operator", "Maintenance and Repair Worker", "Maintenance Equipment Operator", "Maintenance Supervisor", "Maintenance Worker", "Makeup Artists", "Management Analyst", "Manager", "Manager of Air Crew", "Manager of Food Preparation", "Manager of Weapons Specialists", "Manager Tactical Operations", "Manicurists", "Manufactured Building Installer", "Manufacturing Sales Representative", "Mapping Technician", "MARCOM Director", "MARCOM Manager", "Marine Architect", "Marine Cargo Inspector", "Marine Engineer", "Marine Oiler", "Market Research Analyst", "Marketing Manager", "Marketing VP", "Marking Clerk", "Marking Machine Operator", "Marriage and Family Therapist", "Massage Therapist", "Material Movers", "Material Moving Worker", "Materials Engineer", "Materials Inspector", "Materials Scientist", "Mathematical Science Teacher", "Mathematical Scientist", "Mathematical Technician", "Mathematician", "Meat Packer", "Mechanical Door Repairer", "Mechanical Drafter", "Mechanical Engineer", "Mechanical Engineering Technician", "Mechanical Equipment Sales Representative", "Mechanical Inspector", "Media and Communication Worker", "Medical Appliance Technician", "Medical Assistant", "Medical Equipment Preparer", "Medical Equipment Repairer", "Medical Laboratory Technologist", "Medical Records Technician", "Medical Sales Representative", "Medical Scientists", "Medical Secretary", "Medical Technician", "Medical Transcriptionist", "Mental Health Counselor", "Merchandise Displayer OR Window Trimmer", "Metal Fabricator", "Metal Molding Operator", "Metal Pourer and Caster", "Metal Worker", "Metal-Refining Furnace Operator", "Meter Mechanic", "Microbiologist", "Middle School Teacher", "Military Officer", "Milling Machine Operator", "Millwright", "Mine Cutting Machine Operator", "Mining Engineer OR Geological Engineer", "Mining Machine Operator", "Mixing and Blending Machine Operator", "Model Maker", "Mold Maker", "Molder", "Molding and Casting Worker", "Molding Machine Operator", "Motion Picture Projectionist", "Motor Vehicle Inspector", "Motor Vehicle Operator", "Motorboat Mechanic", "Motorboat Operator", "Motorcycle Mechanic", "Movers", "Movie Director oR Theatre Director", "Multi-Media Artist", "Multiple Machine Tool Setter", "Municipal Clerk", "Municipal Court Clerk", "Municipal Fire Fighter", "Municipal Fire Fighting Supervisor", "Museum Conservator", "Music Arranger and Orchestrator", "Music Composer", "Music Director", "Musical Instrument Tuner", "Musician", "Musician OR Singer",
	"Natural Sciences Manager", "Naval Architects", "Network Admin OR Computer Systems Administrator", "Network Systems Analyst", "New Accounts Clerk", "Night Security Guard", "Night Shift", "Nonfarm Animal Caretaker", "Nuclear Engineer", "Nuclear Equipment Operation Technician", "Nuclear Medicine Technologist", "Nuclear Monitoring Technician", "Nuclear Power Reactor Operator", "Nuclear Technician", "Numerical Control Machine Tool Operator", "Numerical Tool Programmer OR Process Control Programmer", "Nursery Manager", "Nursery Worker", "Nursing Aide", "Nursing Instructor", "Nutritionist",
	"Obstetrician", "Occupational Health Safety Specialist", "Occupational Health Safety Technician", "Occupational Therapist", "Occupational Therapist Aide", "Occupational Therapist Assistant", "Office and Administrative Support Worker", "Office Clerk", "Office Machine and Cash Register Servicer", "Office Machine Operator", "Offset Lithographic Press Operator", "Oil and gas Operator", "Oil Service Unit Operator", "Online Marketing Analyst", "Operating Engineer", "Operations Research Analyst", "Ophthalmic Laboratory Technician", "Optical Instrument Assembler", "Opticians", "Optometrist", "Oral Surgeon", "Order Clerk", "Order Filler", "Order Filler OR Stock Clerk", "Organizational Development Manager", "Orthodontist", "Orthotist OR Prosthetist", "Outdoor Power Equipment Mechanic",
	"Packaging Machine Operator", "Packer and Packager", "Painter", "Painter and Illustrator", "Painting Machine Operator", "Pantograph Engraver", "Paper Goods Machine Operator", "Paperhanger", "Paralegal", "Park Naturalist", "Parking Enforcement Worker", "Parking Lot Attendant", "Parts Salesperson", "Paste-Up Worker", "Pastry Chef", "Patrol Officer", "Patternmaker", "Paving Equipment Operator", "Payroll Clerk", "Pediatricians", "Percussion Instrument Repairer", "Personal Care Worker", "Personal Financial Advisor", "Personal Home Care Aide", "Personal Service Worker", "Personal Trainer", "Personnel Recruiter", "Pest Control Worker", "Pesticide Sprayer", "Petroleum Engineer", "Petroleum Pump Operator", "Petroleum Pump System Operator", "Petroleum Technician", "Pewter Caster", "Pharmaceutical Sales Representative", "Pharmacist", "Pharmacy Aide", "Pharmacy Technician", "Philosophy and Religion Teacher", "Photoengraver", "Photoengraving Machine Operator", "Photographer", "Photographic Restorer", "Photographic Developer", "Photographic Process Worker", "Photographic Processing Machine Operator", "Photographic Reproduction Technician", "Physical Scientist", "Physical Therapist", "Physical Therapist Aide", "Physical Therapist Assistant", "Physician", "Physician Assistant", "Physicist", "Physics Teacher", "Pile-Driver Operator", "Pipe Fitter", "Pipefitter", "Pipelayer", "Pipelaying Fitter", "Plant and System Operator", "Plant Scientist", "Plasterer OR Stucco Mason", "Plastic Molding Machine Operator", "Plate Finisher", "Platemaker", "Plating Machine Operator", "Plating Operator", "Plating Operator OR Coating Machine Operator", "Plumber", "Plumber OR Pipefitter OR Steamfitter", "Podiatrist", "Poet OR Lyricist", "Police and Sheriffs Patrol Officer", "Police Detective", "Police Identification OR Records Officer", "Political Science Teacher", "Political Scientist", "Portable Power Tool Repairer", "Postal Clerk", "Postal Service Clerk", "Postal Service Mail Carrier", "Postal Service Mail Sorter", "Postmasters", "Postsecondary Education Administrators", "Postsecondary Teacher", "Potter", "Poultry Cutter", "Power Distributors OR Dispatcher", "Power Generating Plant Operator", "Power Plant Operator", "PR Manager", "Precious Stone Worker", "Precision Aircraft Systems Assemblers", "Precision Devices Inspector", "Precision Dyer", "Precision Etcher and Engraver", "Precision Instrument Repairer", "Precision Lens Grinders and Polisher", "Precision Mold and Pattern Caster", "Precision Pattern and Die Caster", "Precision Printing Worker", "Prepress Technician", "Preschool Education Administrators", "Preschool Teacher", "Press Machine Setter, Operator", "Pressing Machine Operator", "Pressure Vessel Inspector", "Printing Machine Operator", "Printing Press Machine Operator", "Private Detective and Investigator", "Private Household Cook", "Private Sector Executive", "Probation Officers and Correctional Treatment Specialist", "Procurement Clerk", "Producer", "Producers and Director", "Product Management Leader", "Product Promoter", "Product Safety Engineer", "Product Specialist", "Production Control Manager", "Production Helper", "Production Inspector", "Production Laborer", "Production Manager", "Production Planner", "Production Planning", "Production Worker", "Professional Photographer", "Professor", "Program Director", "Project Manager", "Proofreaders and Copy Marker", "Prosthodontist", "Protective Service Worker", "Protective Service Worker", "Psychiatric Aide", "Psychiatric Technician", "Psychiatrist", "Psychologist", "Psychology Teacher", "Public Health Social Worker", "Public Relations Manager", "Public Relations Specialist", "Public Transportation Inspector", "Pump Operators", "Punching Machine Setters", "Purchasing Agent", "Purchasing Manager",
	"Radar Technician", "Radiation Therapist", "Radio and Television Announcer", "Radio Mechanic", "Radio Operator", "Radiologic Technician", "Radiologic Technologist", "Radiologic Technologist and Technician", "Rail Car Repairer", "Rail Transportation Worker", "Rail Yard Engineer", "Railroad Conductors", "Railroad Inspector", "Railroad Switch Operator", "Railroad Yard Worker", "Range Manager", "Real Estate Appraiser", "Real Estate Association Manager", "Real Estate Broker", "Real Estate Sales Agent", "Receptionist and Information Clerk", "Record Clerk", "Recordkeeping Clerk", "Recreation and Fitness Studies Teacher", "Recreation Worker", "Recreational Therapist", "Recreational Vehicle Service Technician", "Recruiter", "Recyclable Material Collector", "Refinery Operator", "Refractory Materials Repairer", "Refrigeration Mechanic", "Registered Nurse", "Rehabilitation Counselor", "Religious Worker", "Rental Clerk", "Reporters OR Correspondent", "Reservation Agent OR Transportation Ticket Agent", "Residential Advisor", "Respiratory Therapist", "Respiratory Therapy Technician", "Restaurant Cook", "Retail Sales person", "Retail Salesperson", "Rigger", "RN", "Rock Splitter", "Rolling Machine Setter", "Roof Bolters Mining", "Roofer", "Rotary Drill Operator", "Rough Carpenter", "Roustabouts",
	"Safety Engineer", "Sailor", "Sales and Related Workers", "Sales Engineer", "Sales Manager", "Sales Person", "Sales Representative", "Sawing Machine Operator", "Sawing Machine Setter", "Sawing Machine Tool Setter", "Scanner Operator", "School Bus Driver", "School Social Worker", "Scientific Photographer", "Screen Printing Machine Operator", "Sculptor", "Secondary School Teacher", "Secretary", "Securities Sales Agent", "Security Guard", "Security Systems Installer OR Fire Alarm Systems Installer", "Segmental Paver", "Self-Enrichment Education Teacher", "Semiconductor Processor", "Separating Machine Operators", "Septic Tank Servicer", "Service Station Attendant", "Set and Exhibit Designer", "Set Designer", "Sewing Machine Operator", "Shampooer", "Shear Machine Set-Up Operator", "Sheet Metal Worker", "Sheriff", "Ship Captain", "Ship Carpenter and Joiner", "Ship Engineer", "Ship Mates", "Ship Pilot", "Shipping and Receiving Clerk", "Shoe and Leather Repairer", "Shoe Machine Operators", "Short Order Cook", "Shuttle Car Operator", "Signal Repairer OR Track Switch Repairer", "Silversmith", "Singer", "Sketch Artist", "Skin Care Specialist", "Slot Key Person", "Social and Human Service Assistant", "Social Media Marketing Manager", "Social Science Research Assistant", "Social Sciences Teacher", "Social Scientists", "Social Service Specialists", "Social Work Teacher", "Social Worker", "Sociologist", "Sociology Teacher", "Software Engineer", "Soil Conservationist", "Soil Scientist", "Soil Scientist OR Plant Scientist", "Solderer", "Soldering Machine Setter", "Sound Engineering Technician", "Space Sciences Teacher", "Special Education Teacher", "Special Force", "Special Forces Officer", "Speech-Language Pathologist", "Sports Book Writer", "Spotters", "Spraying Machine Operator", "Staff Psychologist", "State", "Statement Clerk", "Stationary Engineer", "Stationary Engineer OR Boiler Operator", "Statistical Assistant", "Statistician", "Steel Worker", "Stevedore", "Stock Broker", "Stock Clerk", "Stone Cutter", "Stone Sawyer", "Stonemason", "Stonemason", "Storage Manager OR Distribution Manager", "Streetcar Operator", "Stringed Instrument Repairer and Tuner", "Structural Iron and Steel Worker", "Structural Metal Fabricator", "Substance Abuse Counselor", "Substance Abuse Social Worker", "Substation Maintenance", "Supervisor Correctional Officer", "Supervisor Fire Fighting Worker", "Supervisor of Customer Service", "Supervisor of Police", "Surgeon", "Surgical Technologist", "Survey Researcher", "Surveying and Mapping Technician", "Surveying Technician", "Surveyor", "Sys Admin", "System Administrator",
	"Tailor", "Talent Acquisition Manager", "Talent Director", "Tank Car", "Taper", "Tax Examiner", "Tax Preparer", "Taxi Drivers and Chauffeur", "Teacher", "Teacher Assistant", "Team Assembler", "Technical Director", "Technical Program Manager", "Technical Specialist", "Technical Writer", "Telecommunications Equipment Installer", "Telecommunications Facility Examiner", "Telecommunications Line Installer", "Telemarketer", "Telephone Operator", "Telephone Station Installer and Repairer", "Teller", "Terrazzo Workes and Finisher", "Textile Cutting Machine Operator", "Textile Dyeing Machine Operator", "Textile Knitting Machine Operator", "Textile Machine Operator", "Textile Worker", "Therapist", "Ticket Agent", "Tile Setter OR Marble Setter", "Timing Device Assemblers", "Tire Builder", "Tire Changer", "Title Abstractor", "Title Examiner", "Title Searcher", "Tool and Die Maker", "Tool Set-Up Operator", "Tool Sharpener", "Tour Guide", "Tractor Operator", "Tractor-Trailer Truck Driver", "Traffic Technician", "Train Crew", "Trainer", "Training Manager OR Development Manager", "Transformer Repairer", "Transit Police OR Railroad Police", "Transportation and Material-Moving", "Transportation Attendant", "Transportation Equipment Maintenance", "Transportation Equipment Painters", "Transportation Inspector", "Transportation Manager", "Transportation Worker", "Travel Agent", "Travel Clerk", "Travel Guide", "Tree Trimmer", "Truck Driver", "TSA", "Typesetter", "Typesetting Machine Operator",
	"Umpire and Referee", "Underground Mining", "University", "Upholsterer", "Urban Planner", "User Experience Manager", "User Experience Researcher", "Usher", "Utility Meter Reader",
	"Valve Repairer OR Regulator Repairer", "Vending Machine Servicer", "Veterinarian", "Veterinary Assistant OR Laboratory Animal Caretaker", "Veterinary Technician", "Vice President Of Human Resources", "Vice President Of Marketing", "Video Editor", "Visual Designer", "Vocational Education Teacher",
	"Waiter", "Waitress", "Warehouse", "Washing Equipment Operator", "Waste Treatment Plant Operator", "Watch Repairer", "Weapons Specialists", "Web Developer", "Webmaster", "Welder", "Welder", "Welder and Cutter", "Welder-Fitter", "Welding Machine Tender", "Welding Machine Operator", "Welding Machine Setter", "Welfare Eligibility Clerk", "Well and Core Drill Operator", "Wellhead Pumper", "Wholesale Buyer", "Wind Instrument Repairer", "Woodworker", "Woodworking Machine Operator", "Woodworking Machine Setter", "Word Processors and Typist", "Writer OR Author",
	"Zoologists OR Wildlife Biologist",
}
