rm -rf ./channel-artifacts/*
rm -rf ./crypto-config/*

cryptogen generate --config=./crypto-config.yaml

# 生成创始块文件
echo "---------------- Create genesis.block file BEGIN --------------------"
configtxgen -profile OrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
echo "---------------- Create genesis.block file END --------------------"

# 生成 物流信息channel 文件
echo "---------------- Create TransitOrgschannel.tx file BEGIN -------------------"
configtxgen -profile TransitOrgschannel -outputCreateChannelTx ./channel-artifacts/TransitOrgschannel.tx -channelID TransitOrgschannel
echo "---------------- Create TransitOrgschannel.tx file END -------------------"

# 生成 质监信息channel
echo "---------------- Create SupervOrgchannel.tx file BEGIN -------------------"
configtxgen -profile SupervOrgchannel -outputCreateChannelTx ./channel-artifacts/SupervOrgchannel.tx -channelID SupervOrgchannel
echo "---------------- Create SupervOrgchannel.tx file END -------------------"

# 生成 资金流动channel
echo "---------------- Create FundOrgchannel.tx file BEGIN -------------------"
configtxgen -profile FundOrgchannel -outputCreateChannelTx ./channel-artifacts/FundOrgchannel.tx -channelID FundOrgchannel
echo "---------------- Create FundOrgchannel.tx file END -------------------"


