import ctypes

lib = ctypes.CDLL("./TextConvertor")
# 初始化（會載入字典）
lib.Init()

lib.convert.argtypes = [ctypes.c_char_p]
lib.convert.restype = ctypes.c_char_p

text = "### 卫星图像显示，数百名朝鲜军队可能经海路被运送到俄罗斯参与对乌克兰的战争。这些军队被转移到俄罗斯的杜奈（Dunai）一个隐蔽的军事港口，至少有两艘俄罗斯海军舰艇参与了转移。转移的秘密性表明了避免被发现的努力。这些军队可能被派遣以获得战斗经验并提升其军事能力。朝鲜和俄罗斯之间的关系通过这种合作得到了加强。据报道，用于运输军队的海路已经停止使用，现在改为使用军用飞机进行运输。朝鲜士兵参与冲突凸显了平壤和莫斯科之间日益加深的联系。"

# 轉換函式
result = lib.convert(ctypes.c_char_p(text.encode("utf-8")))

converted_text = result.decode("utf-8")
print(converted_text)