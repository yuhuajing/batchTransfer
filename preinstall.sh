curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source "$HOME/.cargo/env"
rustc --version

wget  https://github.com/ignite/cli/releases/download/v0.27.1/ignite_0.27.1_linux_amd64.tar.gz
tar -zxvf ignite_0.27.1_linux_amd64.tar.gz -C /opt/ignite
sudo mv /opt/ignite/ignite /usr/local/bin/
ignite version

# upgrade ignite version
# rm $(which ignite)
# install the needed ignite packages