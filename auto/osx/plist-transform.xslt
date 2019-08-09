<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" version="1.0">
  <xsl:param name="version"/>
  <xsl:output method="xml" indent="yes"/>
  <xsl:template match="plist-template">
    <plist version="1.0">
      <dict>
        <key>CFBundleDisplayName</key>
        <string><xsl:value-of select="display-name" /></string>
        <key>CFBundleIdentifier</key>
        <string><xsl:value-of select="bundle-identifier" /></string>
        <key>CFBundlePackageType</key>
        <string>APPL</string>
        <key>CFBundleSignature</key>
        <string>cctw</string>
        <key>CFBundleExecutable</key>
        <string><xsl:value-of select="executable" /></string>
        <key>NSHighResolutionCapable</key>
        <true/>
        <key>CFBundleIconFile</key>
        <string><xsl:value-of select="icon-file" /></string>
        <key>CFBundleVersion</key>
        <string><xsl:value-of select="$version" /></string>
      </dict>
    </plist>
  </xsl:template>
</xsl:stylesheet>
