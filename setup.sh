#
# # Go Get It!
#
# Well, here we are again.
# Another installer script.
# We want to make sure that everyone's using the right version of Go, so here it is.
#
# **This script only works for 64-bit Linux machines.**
# Like, say, a campus machine.
#
# You're welcome to set up Go on your own.
# You don't have to use the script.
# You do have to use the correct version of the Go compiler.
#
# I mean, here you are reading it, so you'll probably realize very quickly that it's not hard to do.
# It boils down to the following:
#
# 1. Download Go
# 2. Verify the download
# 3. Unpack it
# 4. Use it.
#

# If anything fails during this process, we want to bail right away.
set -e

# The SHA256 sum for the version of Go that we want.
DOWNLOAD_SUM="508028aac0654e993564b6e2014bf2d4a9751e3b286661b0b0040046cf18028e go1.7.3.linux-amd64.tar.gz"

# ## Download Go.
#
# Get it from Google's servers
echo ""
echo "**Downloading tarball**"
echo ""
wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz

# ## Verify the download
#
# Check its sha256 sum to make sure it downloaded correctly.
# No corruption or other funny business.
echo ""
echo "**Verifying downloaded tarball**"
echo ""
echo $DOWNLOAD_SUM | sha256sum -c -

# ## Unpack it
#
# Unpack the download.
# It'll unpack to a new, local directory named `go`.
echo ""
echo "**Unpacking Go from downloaded tarball**"
echo ""
tar xf go1.7.3.linux-amd64.tar.gz

# Now that it's unpacked, we don't really need the original download anymore.
echo ""
echo "**Removing downloaded tarball**"
echo ""
rm go1.7.3.linux-amd64.tar.gz

# ## Use it
#
# That's it!
# Time to use it.
echo ""
echo "**I'm done!**"
echo ""
echo "You can get to 'go' with the following:"
echo -e "\tGOROOT=\"\$(pwd)/go\" ./go/bin/go --help"
echo ""
echo "Check the version like this:"
echo -e "\tGOROOT=\"\$(pwd)/go\" ./go/bin/go version"
